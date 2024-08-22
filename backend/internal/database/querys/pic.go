package querys

import (
	"context"
	"log"
)

func CanUserSeeImage(userID, fileName string) (bool, error) {
	query := `SELECT privacy_type, post_id, user_id, group_id FROM post WHERE image=$1`
	var privacyType, posterID string
	var postID, groupID *int
	if err := DB.QueryRow(context.Background(), query, fileName).Scan(&privacyType, &postID, &posterID, &groupID); err != nil {
		log.Printf("database: Failed to get image details: %v\n", err)
		return false, err
	}
	if userID == posterID {
		return true, nil
	}
	if privacyType == "private" {
		if IsFollowing(userID, posterID) {
			return true, nil
		}
	} else if privacyType == "almost_private" {
		isAllowed, err := IsAllowed_AlmostPrivate(*postID, userID)
		if err != nil {
			log.Printf("database: Failed to check if user is allowed to see image: %v\n", err)
			return false, err
		}
		if isAllowed {
			return true, nil
		}
	} else if privacyType == "public" {
		return true, nil
	} else if privacyType == "group" {
		if isMember, err := GroupMember(userID, *groupID); err != nil {
			log.Printf("database: Failed to check if user is a group member: %v\n", err)
			return false, err
		} else if isMember {
			return true, nil
		}
	}
	return false, nil
}
