package querys

import (
	"context"
	"fmt"
	"log"
	"social-network/internal/views/websocket/types"
)

func GetPrivateChatUsernames(userID string) ([]string, error) {
	query := `
	SELECT 
			p.user_id,
			u.user_name
	FROM 
			chat c
	INNER JOIN 
			participant p ON p.chat_id = c.chat_id
	INNER JOIN
			"user" u ON p.user_id = u.user_id
	WHERE 
			c.chat_type = 'private'
			AND c.chat_id IN (
					SELECT 
							chat_id 
					FROM 
							participant 
					WHERE 
							user_id = $1
			)
			AND p.user_id != $1
	`
	rows, err := DB.Query(context.Background(), query, userID)
	if err != nil {
		log.Printf("database: Failed to get private chat usernames: %v", err)
		return nil, err
	}
	defer rows.Close()
	var usernames []string
	for rows.Next() {
		var username string
		if err := rows.Scan(&userID, &username); err != nil {
			log.Printf("database: Failed to scan private chat usernames: %v", err)
			return nil, err
		}
		usernames = append(usernames, username)
	}
	if err := rows.Err(); err != nil {
		log.Printf("database: Failed to iterate over private chat usernames: %v", err)
		return nil, err
	}
	return usernames, nil
}

// CREATE TABLE public.chat (
// 	chat_id       SERIAL PRIMARY KEY,
// 	chat_type     public.chat_type NOT NULL,
// 	group_id      INTEGER,
// 	created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE SET NULL
// );

// CREATE TABLE public.participant (
// 	user_id       UUID NOT NULL,
// 	chat_id       INTEGER NOT NULL,
// 	role          public.role_type NOT NULL,
// 	joined_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	PRIMARY KEY (chat_id, user_id),
// 	FOREIGN KEY (chat_id) REFERENCES public.chat (chat_id) ON DELETE CASCADE,
// 	FOREIGN KEY (user_id) REFERENCES public.user (user_id) ON DELETE CASCADE
// );

// CREATE TABLE public.messages (
// 	message_id     SERIAL PRIMARY KEY,
// 	chat_id        INTEGER NOT NULL,
// 	user_id        UUID NOT NULL,
// 	content        TEXT NOT NULL,
// 	created_at     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	updated_at     TIMESTAMP,
// 	FOREIGN KEY (chat_id) REFERENCES public.chat (chat_id) ON DELETE CASCADE,
// 	FOREIGN KEY (user_id) REFERENCES public.user (user_id) ON DELETE CASCADE
// );

// HasPrivateChat checks if there are any chat messages between the two users and returns the chat ID if it exists and the chat type is private
func HasPrivateChat(userID, recipientID string) (int, error) {
	// Check if there are any chat messages between the two users and return chat_id if it exists and the chat type is private
	query := `
	SELECT p.chat_id
	FROM participant p
	JOIN chat c ON p.chat_id = c.chat_id
	WHERE p.user_id IN ($1, $2) AND c.chat_type = 'private'
	GROUP BY p.chat_id
	HAVING COUNT(DISTINCT p.user_id) = 2;
	`

	// Log the query and parameters for debugging
	// log.Printf("Executing query: %s with parameters userID: %s, recipientID: %s", query, userID, recipientID)

	// Execute the query
	rows, err := DB.Query(context.Background(), query, userID, recipientID)
	if err != nil {
		log.Printf("database: Failed to check chat messages: %v", err)
		return 0, err
	}
	defer rows.Close()

	var chatID int
	if rows.Next() {
		if err := rows.Scan(&chatID); err != nil {
			log.Printf("database: Failed to scan chat ID: %v", err)
			return 0, err
		}
		log.Printf("Found chat ID: %d", chatID)
		return chatID, nil
	}

	log.Printf("No private chat found between users %s and %s", userID, recipientID)
	return 0, fmt.Errorf("no private chat found between users %s and %s", userID, recipientID)
}

// CREATE TABLE public.chat (
// 	chat_id       SERIAL PRIMARY KEY,
// 	chat_type     public.chat_type NOT NULL,
// 	group_id      INTEGER,
// 	created_at    TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	FOREIGN KEY (group_id) REFERENCES public.group (group_id) ON DELETE SET NULL
// );

// CREATE TABLE public.group (
// 	group_id     SERIAL PRIMARY KEY,
// 	title        VARCHAR(255),
// 	description  TEXT,
// 	creator_id   UUID NOT NULL,
// 	created_at   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
// 	FOREIGN KEY (creator_id) REFERENCES public.user (user_id)
// );

// CREATE TABLE public.group_member (
//
//	user_id        UUID NOT NULL,
//	group_id       INTEGER  NOT NULL  ,
//	joined_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
//	FOREIGN KEY (user_id) REFERENCES public.user (user_id),
//	FOREIGN KEY (group_id) REFERENCES public.group (group_id)
//
// );

// GetChatIDByGroupID retrieves the chat ID for the group chat
func GetChatIDByGroupID(userID string, groupID int) (int, error) {
	// Query to retrieve the chat ID where the chat type is 'group', and the user is a member of the group
	query := `
	SELECT c.chat_id
	FROM chat c
	JOIN group_member gm ON gm.group_id = c.group_id
	WHERE c.chat_type = 'group' AND c.group_id = $1 AND gm.user_id = $2;
	`

	// Log the query and parameters for debugging
	// log.Printf("Executing query: %s with parameters groupID: %d, userID: %s", query, groupID, userID)

	// Execute the query
	var chatID int
	err := DB.QueryRow(context.Background(), query, groupID, userID).Scan(&chatID)
	if err != nil {
		log.Printf("database: Failed to get chat ID for group: %v", err)
		return 0, err
	}

	return chatID, nil
}

// CreateChat creates a new chat in the database and assigns the chat ID to the users in the participant table
func CreateChat(chatType, userID, recipientID string) (int, error) {
	// Create a new chat in the database and return the chat_id
	query := `
	INSERT INTO public.chat (chat_type) 
	VALUES ($1)
	RETURNING chat_id;
	`

	var chatID int
	err := DB.QueryRow(context.Background(), query, chatType).Scan(&chatID)
	if err != nil {
		log.Printf("database: Failed to create chat: %v", err)
		return 0, err
	}

	// Assign the chat ID to the users in the participant table
	assignQuery := `
	INSERT INTO public.participant (user_id, chat_id, role)
	VALUES ($1, $2, $3), ($4,$5, $6);
	`
	// user_id , chat_id ,role ,member

	_, err = DB.Exec(context.Background(), assignQuery, userID, chatID, "member", recipientID, chatID, "member")
	if err != nil {
		log.Printf("database: Failed to assign chat ID to users: %v", err)
		return 0, err
	}

	return chatID, nil
}

// UpdateChatInDB inserts a new message into the messages table
func UpdateChatInDB(chatID int, message types.Chat, senderID string) error {
	// Insert a new message into the messages table
	query := `
	INSERT INTO public.messages (chat_id, user_id, content)
	VALUES ($1, $2, $3);
	`
	_, err := DB.Exec(context.Background(), query, chatID, senderID, message.Message)
	if err != nil {
		log.Printf("database: Failed to insert message into database: %v", err)
		return err
	}

	return nil
}

// func OpenChatInDB(userID, recipientID int) error {
// 	// Open the chat in the database
// 	query := `
// 		UPDATE chat
// 		SET read = TRUE
// 		WHERE (dm_from = ? AND dm_to = ?) AND read = FALSE
// 	`

// 	_, err := DB.Exec(query, recipientID, userID)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// GetChatMessages retrieves the last 10 messages from the database
func GetChatMessages(chatID int) ([]types.Chat, error) {
	// Get the all messages from the messages table
	query := `
	SELECT m.content, m.user_id, m.created_at
	FROM messages m
	WHERE m.chat_id = $1
	ORDER BY m.created_at DESC;
	`

	// Log the query and parameters for debugging
	// log.Printf("Executing query: %s with parameters chatID: %d", query, chatID)

	// Execute the query
	rows, err := DB.Query(context.Background(), query, chatID)
	if err != nil {
		log.Printf("database: Failed to get messages from database: %v", err)
		return nil, err
	}
	defer rows.Close()

	var messages []types.Chat
	for rows.Next() {
		var message types.Chat
		if err := rows.Scan(&message.Message, &message.Sender, &message.Date); err != nil {
			log.Printf("database: Failed to scan message: %v", err)
			return nil, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
