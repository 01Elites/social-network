package websocket

import (
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"
)

func IsUserConnected(username string) bool {
	_, ok := clients[username]
	return ok
}

func GetClient(userName string) (*types.User, bool) {
	cmutex.Lock()
	defer cmutex.Unlock()
	user, ok := clients[userName]
	return user, ok
}

func SetClientOffline(username string) {
	// Remove the client from the Clients map
	userID := clients[username].ID
	cmutex.Lock()
	delete(clients, username)
	cmutex.Unlock()
	updateFollowersUserList(userID)
}

func SetClientOnline(user *types.User) {
	// Add the client to the Clients map
	cmutex.Lock()
	clients[user.Username] = user
	cmutex.Unlock()
	sendUserList(user)
	updateFollowersUserList(user.ID)
}

func sendUserList(user *types.User) {
	// Get the list of users
	followingUserNames, err := database.GetUserFollowingUserNames(user.ID)
	if err != nil {
		log.Println("Error getting users following:", err)
		return
	}
	listSection := types.Section{
		Name: "Following",
	}
	for _, username := range followingUserNames {
		userProfile, err := database.GetUserProfileByUserName(username)
		if err != nil {
			log.Println("Error getting user profile:", err)
			return
		}
		userDetails := types.UserDetails{
			Username:  userProfile.Username,
			FirstName: userProfile.FirstName,
			LastName:  userProfile.LastName,
		}
		if _, ok := clients[username]; ok {
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
	followers, err := database.GetUserFollowerUserNames(userid)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	for _, followerUsername := range followers {
		if clients[followerUsername] == nil {
			continue
		} else {
			sendUserList(clients[followerUsername])
		}
	}
}
