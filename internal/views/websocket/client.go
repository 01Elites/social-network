package websocket

import (
	"fmt"
	"log"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"
)

func IsUserConnected(username string) bool {
	_, ok := Clients[username]
	return ok
}

func GetClient(userName string) (*types.User, bool) {
	cmutex.Lock()
	defer cmutex.Unlock()
	user, ok := Clients[userName]
	return user, ok
}

func SetClientOffline(user *types.User) {
	// Remove the client from the Clients map
	fmt.Printf("\nSetClientOffline %s\n\n", user.Username)
	cmutex.Lock()
	if Clients[user.Username] == nil {
		cmutex.Unlock()
		return
	}
	userID := Clients[user.Username].ID
	Clients[user.Username].Conn = nil
	delete(Clients, user.Username)
	cmutex.Unlock()
	updateFollowersUserList(userID)
}

func SetClientOnline(user *types.User) {
	// Add the client to the Clients map
	fmt.Printf("\nSetClientOnline %s\n\n", user.Username)
	cmutex.Lock()
	defer cmutex.Unlock()
	Clients[user.Username] = user
	updateFollowersUserList(user.ID)
}

func GetUserList(user *types.User) {
	sendUserList(user)
	dmUserList(user)
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
			Avatar:    userProfile.Avatar,
		}
		if _, ok := Clients[username]; ok {
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
	dms, err := database.GetPrivateChatUsernames(userid)
	if err != nil {
		log.Print("error getting dms:", err)
		return
	}
	for _, followerUsername := range followers {
		if Clients[followerUsername] != nil {
			sendUserList(Clients[followerUsername])
		}
	}
	for _, dmUsername := range dms {
		if Clients[dmUsername] != nil {
			dmUserList(Clients[dmUsername])
		}
	}

}
