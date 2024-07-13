package querys

import (
	"context"
	"log"
	"social-network/internal/models"
	"time"
)

// CreateInvite adds an invitation to the group_invitations table
func CreateInvite(groupID int, senderID string, receiverID string) (int, error) {
	var status string
	var invitationID int
	query := `
	SELECT
		 status,
		 invitation_id
		FROM
			group_invitations
		WHERE
		 (group_id = $1 AND receiver_id = $2)
`
	err := DB.QueryRow(context.Background(), query, groupID, receiverID).Scan(&status, &invitationID)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for invitation: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if status == "pending" {
		log.Printf("invite already made")
		return invitationID, nil
	}
	query = `
    INSERT INTO
        group_invitations (group_id, sender_id, receiver_id)
    VALUES
        ($1, $2, $3)
		RETURNING
				invitation_id`
	err = DB.QueryRow(context.Background(), query, groupID, senderID, receiverID).Scan(&invitationID)
	if err != nil {
		log.Printf("database: Failed to insert invitation into database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	return invitationID, nil
}

// RespondToInvite responds to an invitation that already exists in the group_invitations table
func RespondToInvite(response models.GroupResponse, userID string) error {
	query := `UPDATE group_invitations SET status = $1 WHERE group_id = $2 AND receiver_id = $3 AND status = 'pending'`
	_, err := DB.Exec(context.Background(), query, response.Status, response.GroupID, userID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return err // Return error if failed to update response
	}
	if response.Status == "accepted" {
		query = `INSERT INTO
			group_member (user_id, group_id)
	VALUES
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, userID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return err // Return error if failed to insert group member
		}
	}
	return nil
}

func CheckForGroupRequest(groupID int, senderID string) (bool, error) {
	var requestID int
	query := `
	SELECT
		 request_id
		FROM
			group_requests
		WHERE
		 requester_id = $1 AND status = $2
`
	err := DB.QueryRow(context.Background(), query, senderID, "pending").Scan(&requestID)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for request: %v", err)
		return false, err // Return error if failed to insert post
	}
	if requestID != 0 {
		return true, nil
	}
	return false, nil
}

// CreateRequest adds a request to the group_requests table
func CreateRequest(groupID int, senderID string) (int, string, string, string, error) {
	var creatorUsername string
	var creatorID string
	var groupTitle string
	var requestID int
	query := `
    INSERT INTO
        group_requests (group_id, requester_id)
    VALUES
        ($1, $2)
		RETURNING
			request_id
`
	err := DB.QueryRow(context.Background(), query, groupID, senderID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, "", "","", err // Return error if failed to insert post
	}
	query = `SELECT 
					 public."group".creator_id,
				   public."user".user_name,
					 public."group".title
					 FROM
						"group"
					 INNER JOIN
					 public."user"
					 ON
					 public."user".user_id = public.group.creator_id
					 WHERE
					 group_id = $1`
	err = DB.QueryRow(context.Background(), query, groupID).Scan(&creatorID, &creatorUsername, &groupTitle)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, "", "","", err // Return error if failed to insert post
	}
	// database.AddToNotificationTable()
	// add to notification table for creator
	return requestID, creatorID, creatorUsername, groupTitle, nil
}

// RespondToRequest responds to a request that already exists in the group_requests table
func RespondToRequest(response models.GroupResponse) (int, error) {
	var requestID int
	query := `UPDATE group_requests SET status = $1 WHERE group_id = $2 AND requester_id = $3 AND status = 'pending' RETURNING request_id`
	err := DB.QueryRow(context.Background(), query, response.Status, response.GroupID, response.RequesterID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return 0, err // Return error if failed to insert post
	}
	if response.Status == "accepted" {
		query = `INSERT INTO
			group_member (user_id, group_id)
	VALUES
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, response.RequesterID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return 0,err // Return error if failed to insert post
		}
	}
	return requestID, nil
}

func CancelRequest(GroupID int, userID string) (int, error) {
	var requestID int
	query := `UPDATE group_requests SET status = 'canceled' WHERE requester_id = $1 and group_id = $2 RETURNING request_id`
	err := DB.QueryRow(context.Background(), query, userID, GroupID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return requestID, err // Return error if failed to insert post
	}
	
	return requestID, nil
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

func GetEventOptions(eventID int)([]string, error){
	var options []string
	query:=`SELECT name FROM event_option WHERE event_id = $1`
	rows, err := DB.Query(context.Background(), query, eventID)
	if err != nil {
		log.Print("error getting options")
		return nil, err
	}
	for rows.Next(){
		var option string
		rows.Scan(&option)
		options = append(options, option)
	}
	return options, nil
}

func GetEventDetails(eventID int)(string,int, error){
	var title string
	var groupID int
	query:=`SELECT title, group_id FROM event WHERE event_id = $1`
	err := DB.QueryRow(context.Background(), query, eventID).Scan(&title, &groupID)
	if err != nil {
		log.Print("error scanning title", err)
		return "",0, err
	}
	return title,groupID, nil
}
