package querys

import (
	"context"
	"log"
	"social-network/internal/models"
)

func CreateGroup(userID string, group models.CreateGroup) (int, error){
	query := `
    INSERT INTO 
        "group" (title, description, creator_id) 
    VALUES 
        ($1, $2, $3)
		RETURNING group_id`
		var group_id int
	err := DB.QueryRow(context.Background(),query, group.Title, group.Description, userID).Scan(&group_id)
	if err != nil {
		log.Printf("database: Failed to insert group into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	query = `
	INSERT INTO 
			"group_member" (user_id, group_id) 
	VALUES 
			($1, $2)`
_, err = DB.Exec(context.Background(),query, userID, group_id)
if err != nil {
	log.Printf("database: Failed to insert group into database: %v", err)
	return 0, err // Return error if failed to insert post
}
	return group_id, nil
}

func GetGroupMembers(groupID int) ([]models.User, error) {
	var users []models.User
	query := `SELECT user_id FROM group_member WHERE group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var userID string
		if err = rows.Scan(&userID); err != nil {
			log.Printf("database failed to scan group user: %v\n", err)
			return nil, err
		}
		user, err := GetUserByID(userID)
		if err != nil {
			log.Printf("database failed to get user data: %v\n", err)
			return nil, err
		}
		users = append(users, *user)
	}
	return users, nil
}

func GroupMember(userID string, groupID int) (bool, error) {
	var IsMember int
	query := `SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2`
	err := DB.QueryRow(context.Background(), query, groupID, userID).Scan(&IsMember)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return false, err
	}
	if IsMember == 0{
		return false, nil
	}
	return true, nil
}