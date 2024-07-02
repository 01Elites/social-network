package profile

import (
	"encoding/json"
	"net/http"

	"social-network/internal/database"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func getProfilePosts(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	UserPageID := r.PathValue("id")
	following, err := database.GetUsersFollowingByID(userID)
	if err != nil {
		http.Error(w, "Cannot get user followings", http.StatusInternalServerError)
		return
	}
	posts, err := database.GetUserPosts(userID, UserPageID, following[UserPageID])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}
