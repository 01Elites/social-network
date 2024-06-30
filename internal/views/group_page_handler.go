package views

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"social-network/internal/database"
)

func GetGroupPageHandler(w http.ResponseWriter, r *http.Request){
	groupIDstr := r.PathValue("id")
	groupID, _ := strconv.Atoi(groupIDstr)
	if groupID == 0 {
		http.Error(w, "Invalid group ID", http.StatusBadRequest)
		return
	}
	groupPosts, err := database.GetGroupPosts(groupID)
	if err != nil {
		http.Error(w, "Failed to get group post", http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(groupPosts)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "get post successful")
}