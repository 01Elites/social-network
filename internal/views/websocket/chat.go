package websocket

import (
	"encoding/json"
	"log"
	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"
	"strconv"
	"strings"
	"time"
)

// SendMessage sends a message to a recipient and updates the chat in the database
func SendMessage(RevEvent types.Event, user *types.User) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}

	var message types.Chat
	if err := json.Unmarshal(jsonPayload, &message); err != nil {
		log.Println(err, "error unmarshalling message data")
		return
	}

	// Get the recipient's id by user name
	recipientID, err := database.GetUserIDByUserName(message.Recipient)
	if err != nil {
		log.Println(message.Recipient, "not found in database")
		log.Printf("database: Failed to get recipient: %v", err)
		return
	}

	// cheack if message is empty or white space
	if strings.TrimSpace(message.Message) == "" {
		log.Println("Message is empty")
		return
	}

	// Check if the chat exists
	chatID, _ := database.HasPrivateChat(user.ID, recipientID)

	// Create a new chat if it does not exist
	if chatID == 0 {
		chatID, err = database.CreateChat("private", user.ID, recipientID)
		if err != nil {
			log.Println("Error creating chat:", err)
			return
		}
	}

	// Get the recipient's client from the Clients map and check if it is online
	recipient, online := GetClient(message.Recipient)
	if online {
		// Check if the recipient's chat is opened
		if recipient.ChatOpened == user.Username && recipient.Conn != nil {
			message.Read = true
		}
	}

	// Set the message fields
	message.Sender = user.Username
	message.Date = time.Now().Format("2006-01-02 15:04:05")

	// Update the chat in the database
	err = database.UpdateChatInDB(chatID, message, user.ID)
	if err != nil {
		log.Println("Error updating chat in DataBase:", err)
		return
	}
	var messageResponse struct {
		Messages []types.Chat `json:"messages"`
	}

	messageResponse.Messages = []types.Chat{message}

	// // Convert the message struct to JSON
	// jsonData, err := json.Marshal(messageResponse)
	// if err != nil {
	// 	log.Println(err, "failed to marshal JSON data")
	// 	return
	// }

	// Write JSON data to the WebSocket connection of the user
	sendMessageToWebSocket(user, event.GET_MESSAGES, messageResponse)

	// Send the message to the recipient if they are online and has connection
	if online && recipient.Conn != nil {
		// Write JSON data to the WebSocket connection of the recipient
		sendMessageToWebSocket(recipient, event.GET_MESSAGES, messageResponse)

		// Update the notification field of the recipient in the UserList
		// if !message.Read {
		// 	updateNotification(events.Clients[recipientdb.ID], user.ID, true)
		// }

	}

}

// SendMessageToGroup sends a message to a group and updates the chat in the database
func SendMessageToGroup(RevEvent types.Event, user *types.User) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}

	var message types.Chat
	if err := json.Unmarshal(jsonPayload, &message); err != nil {
		log.Println(err, "error unmarshalling message data")
		return
	}

	// Get the group ID from the message recipient
	groupID, err := strconv.Atoi(message.Recipient)
	if err != nil {
		log.Println(message.Recipient, "not avalid group ID")
		log.Printf("Error converting group ID to int: %v", err)
		return
	}

	// cheack if message is empty or white space
	if strings.TrimSpace(message.Message) == "" {
		log.Println("Message is empty")
		return
	}

	// Check if the chat exists
	chatID, err := database.GetChatIDByGroupID(user.ID, groupID)
	if err != nil {
		log.Println("Error getting chat ID by group ID:", err)
		return
	}

	// if group chat does not exist
	if chatID == 0 {
		log.Println("Error Getting the Group chatid:", err)
	}

	// Set the message fields
	message.Sender = user.Username
	message.Date = time.Now().Format("2006-01-02 15:04:05")

	// Update the chat in the database
	err = database.UpdateChatInDB(chatID, message, user.ID)
	if err != nil {
		log.Println("Error updating chat in DataBase:", err)
		return
	}

	var messageResponse struct {
		Messages []types.Chat `json:"messages"`
	}

	messageResponse.Messages = []types.Chat{message}

	// Convert the message struct to JSON
	jsonData, err := json.Marshal(messageResponse)
	if err != nil {
		log.Println(err, "failed to marshal JSON data")
		return
	}

	// Write JSON data to the WebSocket connection of the user
	sendMessageToWebSocket(user, event.GET_MESSAGES, jsonData)

	// Send the message to the group members if they are online and has
	// connection
	members, _, err := database.GetGroupMembers(user.ID, groupID)
	if err != nil {
		log.Println("Error getting group members:", err)
		return
	}

	for _, member := range members {
		// Get the recipient's client from the Clients map
		recipient, online := GetClient(member.UserName)
		if online && recipient.Conn != nil {
			// Write JSON data to the WebSocket connection of the recipient
			sendMessageToWebSocket(recipient, event.GET_MESSAGES, jsonData)

			// Update the notification field of the recipient in the UserList
			// if !message.Read {
			// 	updateNotification(events.Clients[recipientdb.ID], user.ID, true)
			// }

		}
	}
}

// Function to handle the TYPING event type
func Typing(RevEvent types.Event, user *types.User, IsGroup bool) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}

	var typing types.Typing

	if err := json.Unmarshal(jsonPayload, &typing); err != nil {
		log.Println("Error decoding event typing:", err)
		return
	}

	if typing.Recipient == "" {
		log.Println("Recipient is empty")
		return
	}
	if !IsGroup { // if it is a private chat

		_, err := database.GetUserIDByUserName(typing.Recipient)
		if err != nil {
			log.Println(typing.Recipient, "not found in database")
			log.Printf("database: Failed to get recipient: %v", err)
			return
		}

		// Get the recipient's client from the Clients map
		recipient, online := GetClient(typing.Recipient)
		if online && recipient.Conn != nil {

			typing.Recipient = user.Username

			// Convert the typing struct to JSON
			jsonData, err := json.Marshal(typing)
			if err != nil {
				log.Println(err, "failed to marshal JSON data")
				return
			}
			// Write JSON data to the WebSocket connection of the recipient
			sendMessageToWebSocket(recipient, event.TYPING, jsonData)
		}
	} else { // if it is a group chat
		groupID, err := strconv.Atoi(typing.Recipient)
		if err != nil {
			log.Println(typing.Recipient, "not a valid group ID")
			log.Printf("Error converting group ID to int: %v", err)
			return
		}
		members, _, err := database.GetGroupMembers(user.ID, groupID)
		if err != nil {
			log.Println("Error getting group members:", err)
			return
		}
		var GroupTyping struct {
			Recipient string `json:"recipient"`
			GroupID   string `json:"group_id"`
		}
		GroupTyping.Recipient = user.Username
		GroupTyping.GroupID = typing.Recipient

		// Convert the typing struct to JSON
		jsonData, err := json.Marshal(GroupTyping)
		if err != nil {
			log.Println(err, "failed to marshal JSON data")
			return
		}

		for _, member := range members {
			// Get the recipient's client from the Clients map
			recipient, online := GetClient(member.UserName)
			if online && recipient.Conn != nil {
				// Write JSON data to the WebSocket connection of the recipient
				sendMessageToWebSocket(recipient, event.TYPING, jsonData)
			}
		}
	}
}

// Function to handle the CHAT_OPENED event type
func OpenChat(RevEvent types.Event, user *types.User) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}

	// Unmarshal event payload to get recipient
	var payload struct {
		Recipient string `json:"recipient"`
		IsGroup   bool   `json:"is_group"`
	}
	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		log.Println("Error decoding event payload:", err)
		return
	}

	if payload.Recipient == "" {
		log.Println("Recipient is empty")
		return
	}

	RevEvent.Payload = payload

	// Get and send all messages from the database
	GetMessages(RevEvent, user)
}

func CloseChat(user *types.User) {
	cmutex.Lock()
	Clients[user.Username].ChatOpened = ""
	Clients[user.Username].ChatOpenedIsGroup = false
	cmutex.Unlock()
}

// Function to get messages from the database and send them to the recipient
func GetMessages(RevEvent types.Event, user *types.User) {
	// Convert map to JSON
	jsonPayload, err := json.Marshal(RevEvent.Payload)
	if err != nil {
		log.Println("Error marshaling payload to JSON:", err)
		return
	}
	// Unmarshal event payload to get recipient
	var payload struct {
		Recipient string `json:"recipient"`
		IsGroup   bool   `json:"is_group"`
	}

	if err := json.Unmarshal(jsonPayload, &payload); err != nil {
		log.Println("Error decoding event payload:", err)
		return
	}
	chatID := 0
	if !payload.IsGroup {
		recipientID, err := database.GetUserIDByUserName(payload.Recipient)
		if err != nil {
			log.Println(payload.Recipient, "not found in database")
			log.Printf("database: Failed to get recipient: %v", err)
			return
		}

		cmutex.Lock()
		Clients[user.Username].ChatOpened = payload.Recipient
		Clients[user.Username].ChatOpenedIsGroup = false
		cmutex.Unlock()

		// Check if the chat exists
		chatID, _ = database.HasPrivateChat(user.ID, recipientID)
	} else {
		// from string to int
		groupID, _ := strconv.Atoi(payload.Recipient)
		chatID, err = database.GetChatIDByGroupID(user.ID, groupID)
		if err != nil {
			log.Println(payload.Recipient, "not found in database")
			log.Printf("database: Failed to get recipient: %v", err)
			return
		}

		cmutex.Lock()
		Clients[user.Username].ChatOpened = payload.Recipient
		Clients[user.Username].ChatOpenedIsGroup = true
		cmutex.Unlock()
	}
	if chatID != 0 {
		messages, err := database.GetChatMessages(chatID)
		if err != nil {
			log.Println("Error getting messages from DataBase:", err)

			return
		}

		var messageResponse struct {
			Messages []types.Chat `json:"messages"`
		}

		// messageResponse.Messages = messages

		// if len(messages) != 0 {
		// 	sendMessageToWebSocket(user, event.GET_MESSAGES, messageResponse)
		// }

		if len(messages) != 0 {
			for i := range messages {
				messageResponse.Messages = messages[i : i+1]
				sendMessageToWebSocket(user, event.GET_MESSAGES, messageResponse)
			}
		}
		// updateNotification(Clients[user.ID], recipientID, false)
	}
}
