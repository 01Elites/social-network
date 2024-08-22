package querys

import (
	"context"
	"fmt"
	"log"
)

func CanUserSeeImage(userID, fileName string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) FROM profile WHERE image=$1`
	if err := DB.QueryRow(context.Background(), query, fileName).Scan(&count); err != nil {
		log.Printf("database: Failed to check if user can see image: %v\n", err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	query = `SELECT post_id FROM comment WHERE image=$1`
	var postID *int
	if err := DB.QueryRow(context.Background(), query, fileName).Scan(&postID); err != nil {
		if err.Error() != "no rows in result set" {
			log.Printf("database: Failed to check if user can see image: %v\n", err)
			return false, err
		}
	}
	var privacyType, posterID string
	var groupID *int
	fmt.Println("postID: ", postID)
	if *postID != 0 {
		query = `SELECT privacy_type, user_id, group_id FROM post WHERE post_id=$1`
		if err := DB.QueryRow(context.Background(), query, postID).Scan(&privacyType, &posterID, &groupID); err != nil {
			log.Printf("database: Failed to get image details: %v\n", err)
			return false, err
		}
		fmt.Println("commmeeennnnttt")
	} else {
		query = `SELECT privacy_type, post_id, user_id, group_id FROM post WHERE image=$1`
		if err := DB.QueryRow(context.Background(), query, fileName).Scan(&privacyType, &postID, &posterID, &groupID); err != nil {
			log.Printf("database: Failed to get image details: %v\n", err)
			return false, err
		}
		fmt.Println("posttttt")
	}
	fmt.Println(postID)
	canSee, err := CanSeePostImage(userID, posterID, privacyType, postID, groupID)
	if err != nil {
		log.Printf("database: Failed to check if user can see image: %v\n", err)
		return false, err
	}
	if canSee {
		return true, nil
	}
	return false, nil
}

func CanSeePostImage(userID, posterID, privacyType string, postID, groupID *int) (bool, error) {
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
