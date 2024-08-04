package querys

import (
	"context"
	"log"
	"social-network/internal/models"
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
func RespondToInvite(response models.GroupResponse, userID string) (int, error) {
	var inviteID int
	query := `UPDATE group_invitations SET status = $1 WHERE group_id = $2 AND receiver_id = $3 AND status = 'pending' RETURNING invitation_id`
	err := DB.QueryRow(context.Background(), query, response.Status, response.GroupID, userID).Scan(&inviteID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return 0, err // Return error if failed to update response
	}
	if response.Status == "accepted" {
		query = `INSERT INTO
			group_member (user_id, group_id)
	VALUES
			($1, $2)`
		_, err = DB.Exec(context.Background(), query, userID, response.GroupID)
		if err != nil {
			log.Printf("database: Failed to add group member: %v", err)
			return 0, err // Return error if failed to insert group member
		}
	}
	return inviteID, nil
}

func CheckForGroupInvitation(groupID int, userID string) (models.Requester, error) {
	var invitation models.Requester
	query := `
	SELECT
		 sent_at, first_name, last_name, user_name
		FROM
			group_invitations
		INNER JOIN profile ON public.profile.user_id = public.group_invitations.sender_id
		INNER JOIN "user" USING (user_id)
		WHERE
		 group_id = $1 AND receiver_id = $2 AND status = $3
`
	err := DB.QueryRow(context.Background(), query, groupID, userID, "pending").Scan(&invitation.CreationDate,
		&invitation.User.FirstName, &invitation.User.LastName, &invitation.User.UserName)
	if err != nil && err.Error() != "no rows in result set" {
		log.Printf("database: Failed check for invitation: %v", err)
		return models.Requester{}, err
	}
	if invitation.User.FirstName == "" {
		return models.Requester{}, nil
	}
	return invitation, nil
}

func CheckForGroupRequest(groupID int, senderID string) (bool, error) {
	var requestID int
	query := `
	SELECT
		 request_id
		FROM
			group_requests
		WHERE
		 requester_id = $1 AND group_id = $2 AND status = $3
`
	err := DB.QueryRow(context.Background(), query, senderID, groupID, "pending").Scan(&requestID)
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
func CreateRequest(groupID int, senderID string) (int, string, string, error) {
	var creatorUsername string
	var creatorID string
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
		return 0, "", "", err // Return error if failed to insert post
	}
	query = `SELECT 
					 public."group".creator_id,
				   public."user".user_name
					 FROM
						"group"
					 INNER JOIN
					 public."user"
					 ON
					 public."user".user_id = public.group.creator_id
					 WHERE
					 group_id = $1`
	err = DB.QueryRow(context.Background(), query, groupID).Scan(&creatorID, &creatorUsername)
	if err != nil {
		log.Printf("database: Failed to insert request into database: %v", err)
		return 0, "", "", err // Return error if failed to insert post
	}
	// database.AddToNotificationTable()
	// add to notification table for creator
	return requestID, creatorID, creatorUsername, nil
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
			return 0, err // Return error if failed to insert post
		}
	}
	return requestID, nil
}

func CancelRequest(GroupID int, userID string) (int, error) {
	var requestID int
	query := `UPDATE group_requests SET status = 'canceled' WHERE requester_id = $1 AND group_id = $2 AND status= 'pending' RETURNING request_id`
	err := DB.QueryRow(context.Background(), query, userID, GroupID).Scan(&requestID)
	if err != nil {
		log.Printf("database: Failed to update response in database: %v", err)
		return requestID, err // Return error if failed to insert post
	}

	return requestID, nil
}
