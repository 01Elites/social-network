package querys

import (
	"context"
	"log"
	"time"

	"social-network/internal/models"
)

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
		optionQuery := `SELECT name, option_id FROM event_option WHERE event_id = $1`
		optionRows, err := DB.Query(context.Background(), optionQuery, event.ID)
		if err != nil {
			log.Printf("database failed to query event options: %v\n", err)
			return nil, err
		}
		defer optionRows.Close()

		for optionRows.Next() {
			var option models.Options
			if err = optionRows.Scan(&option.Name, &option.ID); err != nil {
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
			userQuery := `SELECT user_name, first_name, last_name FROM "user" INNER JOIN profile USING (user_id) WHERE user_id = $1`
			var user models.PostFeedProfile
			if err = DB.QueryRow(context.Background(), userQuery, userID).Scan(&user.UserName, &user.FirstName, &user.LastName); err != nil {
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
			if err = DB.QueryRow(context.Background(), optionQuery, OptionID).Scan(&optionName); err != nil {
				log.Printf("database failed to query reponse type for user %v: %v\n", userID, err)
				return nil, err
			}
			for i := range event.Options {
				if event.Options[i].ID == OptionID {
					event.Options[i].Usernames = append(event.Options[i].Usernames, user.UserName)
					event.Options[i].FullNames = append(event.Options[i].FullNames, user.FirstName+" "+user.LastName)
					break
				}
			}
		}
		events = append(events, event)
	}

	if err = rows.Err(); err != nil {
		log.Printf("database rows error: %v\n", err)
		return nil, err
	}
	log.Print("events: ", events)
	return events, nil
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

func CreateEvent(GroupID int, userID string, Title string, Description string, Eventdate time.Time) (int, error) {
	var eventID int
	query := `INSERT INTO event (group_id, creator_id, title, description, event_date) 
	          VALUES ($1, $2, $3, $4, $5) RETURNING event_id`
	err := DB.QueryRow(context.Background(), query, GroupID, userID, Title, Description, Eventdate).Scan(&eventID)
	if err != nil {
		log.Printf("database: Failed to create event: %v", err)
		return 0, err
	}
	return eventID, nil
}

func CreateEventOptions(eventID int, options []string) error {
	query := `INSERT INTO event_option (event_id, name) VALUES ($1, $2)`
	for _, option := range options {
		_, err := DB.Exec(context.Background(), query, eventID, option)
		if err != nil {
			log.Printf("database: Failed to create event options: %v", err)
			return err
		}
	}
	return nil
}

func RespondToEvent(response models.EventResp, userID string) error {
	query := `INSERT INTO user_choice (event_id,user_id,option_id) VALUES ($1,$2,$3)`
	_, err := DB.Exec(context.Background(), query, response.EventID, userID, response.OptionID)
	if err != nil {
		log.Printf("database: Failed to respond to event: %v", err)
		return err
	}
	return nil
}

func CancelEvent(eventID int) error {
	query := `DELETE FROM user_choice WHERE event_id = $1`
	_, err := DB.Exec(context.Background(), query, eventID)
	if err != nil {
		log.Printf("database: Failed to cancel event: %v", err)
		return err
	}
	query = `DELETE FROM event_option WHERE event_id = $1`
	_, err = DB.Exec(context.Background(), query, eventID)
	if err != nil {
		log.Printf("database: Failed to cancel event: %v", err)
		return err
	}
	query = `DELETE FROM event WHERE event_id = $1`
	_, err = DB.Exec(context.Background(), query, eventID)
	if err != nil {
		log.Printf("database: Failed to cancel event: %v", err)
		return err
	}
	return nil
}

func GetEventOptions(eventID int) ([]string, error) {
	var options []string
	query := `SELECT name FROM event_option WHERE event_id = $1`
	rows, err := DB.Query(context.Background(), query, eventID)
	if err != nil {
		log.Print("error getting options")
		return nil, err
	}
	for rows.Next() {
		var option string
		rows.Scan(&option)
		options = append(options, option)
	}
	return options, nil
}

func GetEventDetails(eventID int) (string, int, error) {
	var title string
	var groupID int
	query := `SELECT title, group_id FROM event WHERE event_id = $1`
	err := DB.QueryRow(context.Background(), query, eventID).Scan(&title, &groupID)
	if err != nil {
		log.Print("error scanning title", err)
		return "", 0, err
	}
	return title, groupID, nil
}
