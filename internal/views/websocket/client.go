package websocket

import (
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"

	"github.com/gorilla/websocket"
)

func SetClientOffline(username string) {
	// Remove the client from the Clients map
	cmutex.Lock()
	delete(clients, username)
	cmutex.Unlock()
	updateFollowersUserList(clients[username].ID)
}

func SetClientOnline(conn *websocket.Conn, user *types.User) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[user.Username] = user
	cmutex.Unlock()
	sendUserList(conn, user.ID)
	updateFollowersUserList(user.ID)
}

func sendUserList(conn *websocket.Conn, userID string) {
	// Get the list of users
	followingUserNames, err := database.GetUserFollowingUserNames(userID)
	if err != nil {
		log.Println("Error getting users following:", err)
		return
	}
	listSection := types.Section{
		Name: "Following",
	}
	for _, username := range followingUserNames {
		if user, ok := clients[username]; ok {
			listSection.Users = append(listSection.Users, *user)
		} else {
			// User doesn't exist, create a new offline user
			newUser := types.User{
				Username: username,
				State:    "offline",
				Conn:     nil,
			}
			listSection.Users = append(listSection.Users, newUser)
		}
	}
	// Send the list of users to the client
	sendMessageToWebSocket(conn, event.USERLIST, listSection)
}

func updateFollowersUserList(userid string) {
	followers, err := database.GetUsersFollowees(userid)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	for followerID := range followers {
		followerUsername, err := database.GetUserNameByID(followerID)
		if err != nil {
			log.Print("error getting username:", err)
			return
		}
		if clients[followerUsername] == nil {
			continue
		} else {
			sendUserList(clients[followerUsername].Conn, followerID)
		}
	}
}
