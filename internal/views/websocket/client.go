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
	dmUserList(user)
	updateFollowersUserList(user.ID)
}

func dmUserList(user *types.User) {
	usernames, err := database.GetPrivateChatUsernames(user.ID)
	if err != nil {
		log.Println("Error getting direct message users:", err)
		return
	}
	listSection := buildUserListSection("Direct Messages", usernames)
	sendMessageToWebSocket(user, event.USERLIST, listSection)
}

func sendUserList(user *types.User) {
	usernames, err := database.GetUserFollowingUserNames(user.ID)
	if err != nil {
		log.Println("Error getting following users:", err)
		return
	}
	listSection := buildUserListSection("Following", usernames)
	sendMessageToWebSocket(user, event.USERLIST, listSection)
}

func buildUserListSection(sectionName string, usernames []string) types.Section {
	listSection := types.Section{
		Name: sectionName,
	}
	for _, username := range usernames {
		userProfile, err := database.GetUserProfileByUserName(username)
		if err != nil {
			log.Printf("Error getting user profile for %s: %v", username, err)
			continue
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
	return listSection
}

func updateFollowersUserList(userid string) {
	followers, err := database.GetUserFollowerUserNames(userid)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	for _, followerUsername := range followers {
		if clients[followerUsername] != nil {
			sendUserList(clients[followerUsername])
		}
	}
}
