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
	// Logging to verify the function execution
	fmt.Printf("\nSetClientOffline %s\n\n", user.Username)
	// Locking the user mutex to safely update the user's connections
	user.Mutex.Lock()
	defer user.Mutex.Unlock()

	// Removing the websocket connection from the user's connection map
	delete(user.Conns, conn)

	// Removing the user from the clients map
	delete(clients, user.Username)

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
	dmUserList(user)
	sendUserList(user)
}

func dmUserList(user *types.User) {
	usernames, err := database.GetPrivateChatUsernames(user.ID)
	if err != nil {
		log.Println("Error getting direct message users:", err)
		return
	}
	listSection := buildUserListSection(types.List.DirectMessages, usernames)
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
	listSection := buildUserListSection(types.List.Following, usernames)
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
			fmt.Println("User is online", username)
		} else {
			userDetails.State = types.State.Offline
			fmt.Println("User is offline", username)

		}
		listSection.Users = append(listSection.Users, userDetails)
	}
	return listSection
}

func updateUserInUserList(user *types.User, state string) {
	Payload := types.UserList{
		Type: "update",
		Metadata: types.UserDetails{
			Username: user.Username,
			State:    state,
		},
	}
	followers, err := database.GetUserFollowerUserNames(user.ID)
	if err != nil {
		log.Print("error getting followers:", err)
		return
	}
	dms, err := database.GetPrivateChatUsernames(user.ID)
	if err != nil {
		log.Print("error getting dms:", err)
		return
	}
	// Combine arrays and filter duplicates
	uniqueUsernames := make(map[string]struct{})
	for _, username := range followers {
		uniqueUsernames[username] = struct{}{}
	}
	for _, username := range dms {
		uniqueUsernames[username] = struct{}{}
	}

	// Send messages to all unique clients
	for username := range uniqueUsernames {
		if clients[username] != nil {
			sendMessageToWebSocket(clients[username], event.USERLIST, Payload)
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
