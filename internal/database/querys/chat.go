package querys

import (
	"context"
	"log"
	"social-network/internal/views/websocket/types"
)

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

func HasChat(userID, recipientID string) (bool, error) {
	// Check if there are any chat messages between the two users
	query := `
	SELECT chat_id
	FROM participant
	WHERE user_id IN ($1, $2)
	GROUP BY chat_id
	HAVING COUNT(DISTINCT user_id) = 2;
	`

	// Execute the query
	rows, err := DB.Query(context.Background(), query, userID, recipientID)
	if err != nil {
		log.Printf("database: Failed to check chat messages: %v", err)
		return false, err
	}
	defer rows.Close()

	// Check if there are any rows returned
	hasChat := false
	for rows.Next() {
		hasChat = true
		break
	}

	if err := rows.Err(); err != nil {
		log.Printf("database: Error iterating over rows: %v", err)
		return false, err
	}

	return hasChat, nil
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
func UpdateChatInDB(chatID int, message types.Chat, senderID, recipientID string) error {
	// Insert a new message into the messages table
	query := `
	INSERT INTO public.messages (chat_id, user_id, content, read)
	VALUES ($1, $2, $3, $4);
	`
	_, err := DB.Exec(context.Background(), query, chatID, senderID, message.Message, message.Read)
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
