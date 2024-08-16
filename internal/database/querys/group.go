package querys

import (
	"context"
	"errors"
	"log"
	"time"

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

func CreatGroupChat(groupID int) error {
	query := `INSERT INTO "chat" (group_id, chat_type) VALUES ($1, $2)`
	_, err := DB.Exec(context.Background(), query, groupID, "group")
	if err != nil {
		log.Printf("database: Failed to insert group into database: %v", err)
		return err // Return error if failed to insert group member
	}
	return nil
}

func GetGroupInfo(groupID int) (string, string, error) {
	var title, description string
	query := `SELECT title, description FROM "group" WHERE group_id = $1`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&title, &description)
	if err != nil {
		log.Printf("database failed to get group info: %v\n", err)
		return "", "", err
	}
	return title, description, nil
}

func GetGroupMembers(loggedUser string, groupID int) ([]models.UserLiteInfo, []string, error) {
	var users []models.UserLiteInfo
	var userids []string
	query := `SELECT user_id, user_name, first_name, last_name, image , privacy FROM group_member 
						INNER JOIN profile USING (user_id)
						INNER JOIN "user" USING (user_id)
						WHERE group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan group user1: %v\n", err)
		return nil, nil, err
	}
	for rows.Next() {
		var user models.UserLiteInfo
		var userid string
		if err = rows.Scan(&userid, &user.UserName, &user.FirstName, &user.LastName, &user.Avatar, &user.Privacy); err != nil {
			log.Printf("database failed to scan group user2: %v\n", err)
			return nil, nil, err
		}
		user.Status = GetFollowStatus(loggedUser, userid)
		users = append(users, user)
		userids = append(userids, userid)
	}
	return users, userids, nil
}

func GroupMember(userID string, groupID int) (bool, error) {
	var IsMember int
	query := `SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2`
	err := DB.QueryRow(context.Background(), query, groupID, userID).Scan(&IsMember)
	if err != nil {
		log.Printf("database failed to scan group user11: %v\n", err)
		return false, err
	}
	if IsMember == 0 {
		return false, nil
	}
	return true, nil
}

func CheckGroupID(groupID int) bool {
	var groupExists int
	query := `SELECT COUNT(*) FROM "group" WHERE group_id = $1`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&groupExists)
	if err != nil {
		log.Printf("database failed to scan group user12: %v\n", err)
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

func GetCreatorProfile(groupID int) (models.PostFeedProfile, error) {
	var creator models.PostFeedProfile
	query := `SELECT user_name, first_name, last_name, image FROM profile INNER JOIN "user" USING (user_id) WHERE user_id = (SELECT creator_id FROM "group" WHERE group_id = $1)`
	err := DB.QueryRow(context.Background(), query, groupID).Scan(&creator.UserName, &creator.FirstName, &creator.LastName, &creator.Avatar)
	if err != nil {
		log.Printf("database failed to get creator profile: %v\n", err)
		return models.PostFeedProfile{}, err
	}
	return creator, nil
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

func getGroupFromRequest(requestID int) (int, string, string, string, string, error) {
	var groupID int
	var count int
	var groupTitle string
	var creator_id string
	var requested_at time.Time
	var requesterID string
	query := `SELECT
						group_id,
						title,
						creator_id,
						requester_id,
						requested_at
						FROM
						group_requests
						INNER JOIN
						"group"	USING	(group_id)
						WHERE
						request_id = $1 AND status = 'pending'
						`
	err := DB.QueryRow(context.Background(), query, requestID).Scan(&groupID, &groupTitle, &creator_id,&requesterID, &requested_at)
	if err != nil {
		if err == errors.New("no rows in result set") {
		log.Printf("database failed to scan group user3: %v\n", err)
		}
		return 0, "", "", "", "", err
	}
	query = `SELECT COUNT(*) FROM group_member WHERE group_id = $1 AND user_id = $2`
	err = DB.QueryRow(context.Background(), query, groupID, requesterID).Scan(&count)
	if err != nil {
		log.Printf("database failed to scan group member")
		return 0, "", "", "", "", err
	}
	if count != 0 {
		log.Print("user is already a member of the group")
		return 0, "", "", "", "", errors.New("UserAlreadyMember")
	}

	return groupID, groupTitle, creator_id, requesterID, requested_at.String(),  nil
}

func getGroupFromInvitation(invitationID int) (string, int, string, models.Requester, error) {
	var groupID int
	var groupTitle string
	var invitedUser string
	var senderID string
	var sentAt time.Time
	query := `SELECT
						receiver_id,
						group_id,
						title,
						sender_id,
						sent_at
						FROM
						group_invitations
						INNER JOIN
						"group"	USING	(group_id)
						WHERE
						invitation_id = $1 AND status = 'pending'
						`
	err := DB.QueryRow(context.Background(), query, invitationID).Scan(&invitedUser, &groupID, &groupTitle, &senderID, &sentAt)
	if err != nil {
		if err == errors.New("no rows in result set") {
			log.Printf("database failed to scan group user4: %v\n", err)
		}
		return "", 0, "", models.Requester{}, err
	}
	inviter, err := GetUserPostFeedProfile(senderID)
	if err != nil {
		log.Printf("database failed to get user profile: %v\n", err)
		return "", 0, "", models.Requester{}, err
	}
	inviteData := models.Requester{
		User: *inviter,
		CreationDate: sentAt.String(),
	} 
	return invitedUser, groupID, groupTitle, inviteData, nil
}

func GetGroupRequests(groupID int) ([]models.Requester, error) {
	var requesters []models.Requester
	query := `SELECT user_name,
									 requested_at,
									 first_name,
									 last_name,
									 image
									  FROM group_requests
										INNER JOIN profile ON public.profile.user_id = public.group_requests.requester_id
										INNER JOIN "user" USING (user_id)
										WHERE group_id = $1 AND status = 'pending'`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan group user5: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var requester models.Requester
		var created_at time.Time
		if err = rows.Scan(&requester.User.UserName,
			&created_at,
			&requester.User.FirstName,
			&requester.User.LastName,
			&requester.User.Avatar); err != nil {
			log.Printf("database failed to scan group user6: %v\n", err)
			return nil, err
		}
		requester.CreationDate = created_at.String()
		requesters = append(requesters, requester)
	}
	return requesters, nil
}

func GetAllGroups() {

}

// get the id of all the groups in database
func GetAllGroupIDs() ([]int, error) {
	var groupIDs []int
	query := `SELECT group_id FROM "group"`
	rows, err := DB.Query(context.Background(), query)
	if err != nil {
		log.Printf("database failed to scan group user7: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var groupID int
		if err = rows.Scan(&groupID); err != nil {
			log.Printf("database failed to scan group user8: %v\n", err)
			return nil, err
		}
		groupIDs = append(groupIDs, groupID)
	}
	return groupIDs, nil
}

func GetGroupFeedInfo(groupID int, userID string) (models.GroupFeed, error) {
	var group models.GroupFeed
	group.ID = groupID
	// group.Title, group.Description, group.Creator, group.Requesters, group.Events, group.IsCreator, group.RequestMade = database.GetGroupInfo(groupID)

	username, err := GetUserNameByID(userID)
	if err != nil {
		return models.GroupFeed{}, err
	}

	group.Title, group.Description, err = GetGroupInfo(group.ID)
	if err != nil {
		return models.GroupFeed{}, err
	}

	group.Creator, err = GetCreatorProfile(group.ID)
	if err != nil {
		return models.GroupFeed{}, err

	}

	if group.Creator.UserName == username {
		group.IsCreator = true
	}

	if group.IsMember, err = GroupMember(userID, group.ID); err != nil {
		return models.GroupFeed{}, err
	}

	if !group.IsMember {
		group.RequestMade, err = CheckForGroupRequest(group.ID, userID)
		if err != nil {
			return models.GroupFeed{}, err
		}
	}

	return group, nil
}
