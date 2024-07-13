package querys

import (
	"context"
	"fmt"
	"log"

	"social-network/internal/models"
)

func CreateGroup(userID string, group models.CreateGroup) (int, error) {
	query := `
    INSERT INTO
        "group" (title, description, creator_id)
    VALUES
        ($1, $2, $3)
		RETURNING group_id`
	var group_id int
	err := DB.QueryRow(context.Background(), query, group.Title, group.Description, userID).Scan(&group_id)
	if err != nil {
		log.Printf("database: Failed to insert group into database: %v", err)
		return 0, err // Return error if failed to insert group
	}
	query = `
	INSERT INTO
			"group_member" (user_id, group_id)
	VALUES
			($1, $2)`
	_, err = DB.Exec(context.Background(), query, userID, group_id)
	if err != nil {
		log.Printf("database: Failed to insert group into database: %v", err)
		return 0, err // Return error if failed to insert group member
	}
	return group_id, nil
}

func GetGroupMembers(groupID int) ([]string, []string, error) {
	var users []string
	var userIDs []string
	query := `SELECT user_id FROM group_member WHERE group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return nil, nil, err
	}
	for rows.Next() {
		var userID string
		if err = rows.Scan(&userID); err != nil {
			log.Printf("database failed to scan group user: %v\n", err)
			return nil, nil, err
		}
		username, err := GetUserNameByID(userID)
		if err != nil {
			log.Printf("database failed to get user data: %v\n", err)
			return nil, nil, err
		}
		userIDs = append(userIDs, userID)
		users = append(users, username)
	}
	return users, userIDs, nil
}

func GroupMember(userID string, groupID int) (bool, error) {
	var IsMember int
	query := `SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2`
	err := DB.QueryRow(context.Background(), query, groupID, userID).Scan(&IsMember)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return false, err
	}
	if IsMember == 0 {
		return false, nil
	}
	return true, nil
}

func EventCreator(userID string, eventID int) (bool, error) {
	var creatorID string
	query := `SELECT creator_id FROM "group" WHERE group_id = (SELECT group_id FROM event WHERE event_id = $1)`
	err := DB.QueryRow(context.Background(), query, eventID).Scan(&creatorID)
	if err != nil {
		log.Printf("database failed to scan event creator: %v\n", err)
		return false, err
	}
	if creatorID != userID {
		return false, nil
	}
	return true, nil
}

func CheckGroupID(groupID int) bool {
	var groupExists int
	query := `SELECT COUNT(*) FROM "group" WHERE group_id = $1`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&groupExists)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return false
	}
	if groupExists == 0 {
		return false
	}
	return true
}

func GetGroupTitle(groupID int) (string, error) {
	var groupTitle string
	query := `SELECT title FROM "group" WHERE group_id = $1`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&groupTitle)
	if err != nil {
		log.Printf("database failed to get group title: %v\n", err)
		return "", err
	}
	return groupTitle, nil
}

func GetGroupCreatorID(groupID int) (string, error) {
	var creatorID string
	query := `SELECT creator_id FROM "group" WHERE group_id = $1`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&creatorID)
	if err != nil {
		log.Printf("database failed to scan group creator: %v\n", err)
		return "", err
	}
	return creatorID, nil
}

func CheckEventID(eventID int) int {
	var groupId int
	query := `SELECT group_id FROM event WHERE event_id = $1`
	err := DB.QueryRow(context.Background(), query, eventID).Scan(&groupId)
	if err != nil {
		log.Printf("database failed to scan event: %v\n", err)
		return 0
	}
	if groupId == 0 {
		return 0
	}
	return groupId
}

func LeaveGroup(userID string, groupID int) error {
	query := `DELETE FROM group_member WHERE group_id = $1 AND user_id = $2`
	_, err := DB.Exec(context.Background(), query, groupID, userID)
	if err != nil {
		log.Printf("database failed to delete group user: %v\n", err)
		return err
	}
	return nil
}

func GetGroupEvents(groupID int) ([]models.Event, error) {
	var events []models.Event
	query := `SELECT event_id, title, description, event_date FROM event WHERE group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to query group events: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var event models.Event
		if err = rows.Scan(&event.ID, &event.Title, &event.Description, &event.EventTime); err != nil {
			log.Printf("database failed to scan event: %v\n", err)
			return nil, err
		}

		// Get options for the event
		optionQuery := `SELECT name FROM event_option WHERE event_id = $1`
		optionRows, err := DB.Query(context.Background(), optionQuery, event.ID)
		if err != nil {
			log.Printf("database failed to query event options: %v\n", err)
			return nil, err
		}
		defer optionRows.Close()

		for optionRows.Next() {
			var option string
			if err = optionRows.Scan(&option); err != nil {
				log.Printf("database failed to scan event option: %v\n", err)
				return nil, err
			}
			event.Options = append(event.Options, option)
		}

		// Get responded users for the event
		responseQuery := `SELECT user_id FROM user_choice WHERE event_id = $1`
		responseRows, err := DB.Query(context.Background(), responseQuery, event.ID)
		if err != nil {
			log.Printf("database failed to query event responses: %v\n", err)
			return nil, err
		}
		defer responseRows.Close()

		for responseRows.Next() {
			var userID string
			if err = responseRows.Scan(&userID); err != nil {
				log.Printf("database failed to scan responded user ID: %v\n", err)
				return nil, err
			}

			// Get username for each responded user
			userQuery := `SELECT user_name FROM "user" WHERE user_id = $1`
			var username string
			if err = DB.QueryRow(context.Background(), userQuery, userID).Scan(&username); err != nil {
				log.Printf("database failed to query user_name: %v\n", err)
				return nil, err
			}
			optionIdQuery := `SELECT option_id FROM user_choice WHERE event_id = $1 and user_id = $2`
			var OptionID int
			if err = DB.QueryRow(context.Background(), optionIdQuery, event.ID, userID).Scan(&OptionID); err != nil {
				log.Printf("database failed to query option_id for user %v: %v\n", userID, err)
				return nil, err
			}
			optionQuery := `SELECT name FROM event_option WHERE option_id = $1`
			var optionName string
			fmt.Println(OptionID)
			if err = DB.QueryRow(context.Background(), optionQuery, OptionID).Scan(&optionName); err != nil {
				log.Printf("database failed to query reponse type for user %v: %v\n", userID, err)
				return nil, err
			}
			event.RespondedUsers = append(event.RespondedUsers, username, optionName)
		}

		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Printf("database rows error: %v\n", err)
		return nil, err
	}

	return events, nil
}

func getGroupFromRequest(requestID int) (int, string, error) {
	var groupID int
	var groupTitle string
	query := `SELECT
						group_id,
						title
						FROM
						group_requests
						INNER JOIN
						"group"	USING	(group_id)
						WHERE
						request_id = $1
						`
	err := DB.QueryRow(context.Background(), query, requestID).Scan(&groupID,&groupTitle)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return 0, "", err
	}
	return groupID, groupTitle, nil
}
