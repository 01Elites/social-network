package websocket

import (
	"fmt"
	"log"
	"sync"

	database "social-network/internal/database/querys"
	"social-network/internal/views/websocket/types"
	"social-network/internal/views/websocket/types/event"

	"github.com/gorilla/websocket"
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

func SetClientOffline(user *types.User, conn *websocket.Conn) {
	// Remove the client from the Clients map
	fmt.Printf("\nSetClientOffline %s\n\n", user.Username)
	user.Mutex.Lock()
	defer user.Mutex.Unlock()
	if _, exists := user.Conns[conn]; exists {
		delete(user.Conns, conn)
	}

	if len(user.Conns) == 0 {
		cmutex.Lock()
		defer cmutex.Unlock()
		if _, ok := clients[user.Username]; ok {
			delete(clients, user.Username)
		}
	}
}

func SetClientOnline(userID string, conn *websocket.Conn) (*types.User, error) {
	// Retrieve the username from the database
	username, err := database.GetUserNameByID(userID)
	if err != nil {
		log.Println("Error getting user name:", err)
		return nil, err
	}
	fmt.Printf("\nSetClientOnline %s\n\n", username)
	cmutex.Lock()
	defer cmutex.Unlock()

	user, ok := clients[username]
	if !ok {
		user = &types.User{
			ID:       userID,
			Username: username,
			Conns:    make(map[*websocket.Conn]bool),
			Mutex:    &sync.Mutex{},
		}
		clients[username] = user
	}
	user.Mutex.Lock()
	defer user.Mutex.Unlock()
	user.Conns[conn] = true
	return user, nil
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
	payload := types.UserList{
		Type:     "init",
		Metadata: listSection,
	}
	sendMessageToWebSocket(user, event.USERLIST, payload)
}

func sendUserList(user *types.User) {
	usernames, err := database.GetUserFollowingUserNames(user.ID)
	if err != nil {
		log.Println("Error getting following users:", err)
		return
	}
	listSection := buildUserListSection("Following", usernames)
	payload := types.UserList{
		Type:     "init",
		Metadata: listSection,
	}
	sendMessageToWebSocket(user, event.USERLIST, payload)
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
		if _, ok := clients[username]; ok {
			userDetails.State = types.State.Online
		} else {
			userDetails.State = types.State.Offline
		}
		listSection.Users = append(listSection.Users, userDetails)
	}
	return listSection
}

func updateUserInUserList(user *types.User, state string) {
	followers, err := database.GetUserFollowerUserNames(user.ID)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	Payload := types.UserList{
		Type: "update",
		Metadata: types.UserDetails{
			Username: user.Username,
			State:    state,
		},
	}
	for _, followerUsername := range followers {
		if clients[followerUsername] != nil {
			sendMessageToWebSocket(clients[followerUsername], event.USERLIST, Payload)
		}
	}
}

func AddUserToUserList(toUserID string, userID string, listName string) {
	toUser, err := database.GetUserNameByID(toUserID)
	if err != nil {
		log.Println("Error getting user name:", err)
		return
	}
	user, err := database.GetUserProfile(userID)
	if err != nil {
		log.Println("Error getting user profile:", err)
		return
	}
	userDetails := types.UserDetails{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Username:  user.Username,
		State:     types.State.Offline,
		Avatar:    user.Avatar,
	}
	if clients[user.Username] != nil {
		userDetails.State = types.State.Online
	}
	metadata := map[string]interface{}{
		"name": listName,
		"user": userDetails,
	}
	Payload := types.UserList{
		Type:     "add",
		Metadata: metadata,
	}
	if clients[toUser] != nil {
		sendMessageToWebSocket(clients[toUser], event.USERLIST, Payload)
	}
}
