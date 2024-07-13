package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	database "social-network/internal/database/querys"
	"time"

	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"
	"strings"

	"github.com/gorilla/websocket"
)

func ProcessEvents(user *types.User) {
	defer func() {
		// Remove the client from the Clients map when the connection is closed
		user.Conn.Close()
		SetClientOffline(user.Username)
	}()

	for {
		// Read message from WebSocket connection
		messageType, rawMessage, err := user.Conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message from WebSocket:", err)
			return
		}

		// if the messageType is not string will skip this message
		if messageType != websocket.TextMessage {
			continue
		}

		fmt.Printf("Received message from user %s: %s\n", user.Username, string(rawMessage))
		// Deserialize the message into the Event struct
		var message types.Event
		err = json.Unmarshal(rawMessage, &message)
		if err != nil {
			log.Println("Error unmarshalling JSON message into Event struct:", err)
			return
		}
		fmt.Println("message: ", message)

		// Handle the event based on its type
		switch message.Type {
		case event.SEND_MESSAGE:
			// Call function for NOTIFICATION
			log.Println("SEND_MESSAGE")
			SendMessage(message, user)
		case event.TYPING:
			// Call function for TYPING
			// Typing(event, user)
		case event.CHAT_OPENED:
			// OpenChat(event, user)
		case event.CHAT_CLOSED:
			// CloseChat(user)
		case event.GET_MESSAGES:
			// GetMessages(event, user)
		case event.GET_NOTIFICATIONS:
			// GetNotifications(event, user)
		default:
			log.Println("Unknown event type:", message.Type)
		}
	}
}

/// todo:
// 1. Implement the Typing function
// 2. Implement the OpenChat function
// 3. Implement the CloseChat function
// 4. Implement the GetMessages function
// 5. Implement the GetNotifications function
// 6. get the list of users {online, offline,Dm's } and send it to the client
// 7. get the list of groups {online, offline, messages} and send it to the client

// to do with the notifications
// 1. Implement the notification function {send notification to the user by websocket}
// 2. notifications include:
// 	- new comment on the my post {post_id, commenter , comment}
// 	- new like on the my post {post_id, liker}
// 	- new friend request "follow request" {follower} to be [accepted or rejected]
// 	- new message in DM {sender, message}
// 	- new message in the group chat {sender, message , group_id}
//  - invitation to the group
//  - new reuest to join the group {which the user is the admin}
//  - new event in the group {which the user is the admin} MAYBE !!!

// required notification types:
// follow request
// group invitation
// request to join group
// event in the group

// SendMessage function
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
	recipetsID, err := database.GetUserIDByUserName(message.Recipient)
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
	chatID, _ := database.HasPrivateChat(user.ID, recipetsID)

	// Create a new chat if it does not exist
	if chatID == 0 {
		chatID, err = database.CreateChat("private", user.ID, recipetsID)
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
	err = database.UpdateChatInDB(chatID, message, user.ID, recipetsID)
	if err != nil {
		log.Println("Error updating chat in DataBase:", err)
		return
	}
	var messageResponse struct {
		// Index    int          `json:"index"`
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
	sendMessageToWebSocket(user.Conn, event.GET_MESSAGES, jsonData)

	// Send the message to the recipient if they are online and has connection
	if online && recipient.Conn != nil {
		// Write JSON data to the WebSocket connection of the recipient
		sendMessageToWebSocket(recipient.Conn, event.GET_MESSAGES, jsonData)

		// Update the notification field of the recipient in the UserList
		if !message.Read {
			// updateNotification(events.Clients[recipientdb.ID], user.ID, true)
		}

	}

}

func GetClient(userName string) (*types.User, bool) {
	cmutex.Lock()
	defer cmutex.Unlock()
	user, ok := clients[userName]
	return user, ok
}

// func TransferToDMs(user, recipient *events.Client) {
// 	findAndTransferUserToDMs(user, recipient)
// 	SendUsersListToClient(user.UserID)
// 	findAndTransferUserToDMs(recipient, user)
// 	SendUsersListToClient(recipient.UserID)
// }

// func ReorderList(user, recipient *events.Client) {
// 	SetAtTheTop(user, recipient)
// 	SetAtTheTop(recipient, user)
// 	SendUsersListToClient(user.UserID)
// 	SendUsersListToClient(recipient.UserID)
// }

// func SetAtTheTop(user, recipient *events.Client) {
// 	for i, section := range user.UserList {
// 		if section.Name == "DMs" {
// 			for j, userNot := range section.Users {
// 				if userNot.Client.Details.UserName == recipient.Details.UserName {
// 					// Remove the user from the sender's user list
// 					user.UserList[i].Users = append(user.UserList[i].Users[:j], user.UserList[i].Users[j+1:]...)
// 					user.UserList[i].Users = append([]events.UserNotification{userNot}, user.UserList[i].Users...)
// 				}
// 			}
// 		}
// 	}
// }

// func findAndTransferUserToDMs(user, recipient *events.Client) {
// 	for i, section := range user.UserList {
// 		if section.Name == "Users" {
// 			for j, userNot := range section.Users {
// 				if userNot.Client.Details.UserName == recipient.Details.UserName {
// 					// Remove the user from the sender's user list
// 					user.UserList[i].Users = append(user.UserList[i].Users[:j], user.UserList[i].Users[j+1:]...)
// 					events.Clients[user.UserID].UserList[0].Users = append(events.Clients[user.UserID].UserList[0].Users, userNot)
// 				}
// 			}
// 		}
// 	}
// }

// /********************** send users list **************************/
// func SendUsersList() {
// 	// send the new list to all clients
// 	for _, client := range events.Clients {
// 		if client.Conn != nil {
// 			// Convert the user list to JSON
// 			jsonData, err := json.Marshal(client.UserList)
// 			if err != nil {
// 				log.Println("Error marshalling user list to JSON:", err)
// 				return
// 			}
// 			// Write JSON data to the WebSocket connection
// 			sendMessageToWebSocket(client.Conn, "USERS", jsonData)
// 		}
// 	}
// }

// // SendUsersListToClient sends the user list to a specific client
// func SendUsersListToClient(userID int) {
// 	// send the new list to all clients
// 	client := events.Clients[userID]
// 	if client.Conn == nil {
// 		// log.Println("Client connection is nil for user:", client.Details.UserName)
// 		return
// 	}
// 	// Convert the user list to JSON
// 	jsonData, err := json.Marshal(client.UserList)
// 	if err != nil {
// 		log.Println("Error marshalling user list to JSON:", err)
// 		return
// 	}
// 	// Write JSON data to the WebSocket connection
// 	sendMessageToWebSocket(client.Conn, "USERS", jsonData)
// }
