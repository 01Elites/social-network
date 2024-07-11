package websocket

import (
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"
)

func SetClientOffline(userID string) {
	// Remove the client from the Clients map
	cmutex.Lock()
	delete(clients, userID)
	cmutex.Unlock()
	updateFollowersUserList(userID)
}

func SetClientOnline(user *types.User) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[user.ID] = user
	cmutex.Unlock()
	sendUserList(user)
	updateFollowersUserList(user.ID)
}

func sendUserList(user *types.User) {
	// Get the list of users
	followingIDs, err := database.GetUsersFollowingByID(user.ID)
	if err != nil {
		log.Println("Error getting users following:", err)
		return
	}
	listSection := types.Section{
		Name: "Following",
	}
	for userID := range followingIDs {
		userProfile, err := database.GetUserProfile(userID)
		if err != nil {
			log.Println("Error getting user profile:", err)
			return
		}
		userDetails := types.UserDetails{
			Username:  userProfile.Username,
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
		}
		if _, ok := clients[userID]; ok {
			userDetails.State = "online"
		} else {
			userDetails.State = "offline"
		}
		listSection.Users = append(listSection.Users, userDetails)
	}
	// // Send the list of users to the client
	sendMessageToWebSocket(user.Conn, event.USERLIST, listSection)
}

func updateFollowersUserList(userid string) {
	followers, err := database.GetUsersFollowees(userid)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	for followerID := range followers {
		if clients[followerID] == nil {
			continue
		} else {
			sendUserList(clients[followerID])
		}
	}
}
