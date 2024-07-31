package querys

import (
	"context"
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

func GetGroupMembers(groupID int) ([]models.PostFeedProfile, []string, error) {
	var users []models.PostFeedProfile
	var userids []string
	query := `SELECT user_id, user_name, first_name, last_name, image FROM group_member 
						INNER JOIN profile USING (user_id)
						INNER JOIN "user" USING (user_id)
						WHERE group_id = $1`
	rows, err := DB.Query(context.Background(), query, groupID)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return nil, nil, err
	}
	for rows.Next() {
		var user models.PostFeedProfile
		var userid string
		if err = rows.Scan(&userid, &user.UserName, &user.FirstName, &user.LastName, &user.Avatar); err != nil {
			log.Printf("database failed to scan group user: %v\n", err)
			return nil,nil, err
		}
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
		log.Printf("database failed to scan group user: %v\n", err)
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

func getGroupFromRequest(requestID int) (int, string, string, error) {
	var groupID int
	var groupTitle string
	var creator_id string
	query := `SELECT
						group_id,
						title,
						creator_id
						FROM
						group_requests
						INNER JOIN
						"group"	USING	(group_id)
						WHERE
						request_id = $1
						`
	err := DB.QueryRow(context.Background(), query, requestID).Scan(&groupID, &groupTitle, &creator_id)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return 0, "", "", err
	}
	return groupID, groupTitle, creator_id, nil
}

func getGroupFromInvitation(invitationID int) (string, int, string, error) {
	var groupID int
	var groupTitle string
	var invitedUser string
	query := `SELECT
						receiver_id,
						group_id,
						title
						FROM
						group_invitations
						INNER JOIN
						"group"	USING	(group_id)
						WHERE
						invitation_id = $1
						`
	err := DB.QueryRow(context.Background(), query, invitationID).Scan(&invitedUser, &groupID, &groupTitle)
	if err != nil {
		log.Printf("database failed to scan group user: %v\n", err)
		return "", 0, "", err
	}
	return invitedUser, groupID, groupTitle, nil
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
		log.Printf("database failed to scan group user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var requester models.Requester
		if err = rows.Scan(&requester.User.UserName,
			&requester.CreationDate,
			&requester.User.FirstName,
			&requester.User.LastName,
			&requester.User.Avatar); err != nil {
			log.Printf("database failed to scan group user: %v\n", err)
			return nil, err
		}
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
		log.Printf("database failed to scan group user: %v\n", err)
		return nil, err
	}
	for rows.Next() {
		var groupID int
		if err = rows.Scan(&groupID); err != nil {
			log.Printf("database failed to scan group user: %v\n", err)
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

	group.Members,_, err = GetGroupMembers(group.ID)
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

	return group, nil
}
