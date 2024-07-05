package notifications

import (
	"encoding/json"
	"net/http"
	database "social-network/internal/database/querys"
	"social-network/internal/models"
	"social-network/internal/views/middleware"
)

func CreateFollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var request models.Request
	request.SenderID = userID
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	request.ID, err = database.CreateFollowRequest(request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	requestIDjson := struct {
		ID int `json:"id"`
	}{
		ID: request.ID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(requestIDjson)
}

func RespondToFollowHandler(w http.ResponseWriter, r *http.Request) {
	userID, ok := r.Context().Value(middleware.UserIDKey).(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusInternalServerError)
		return
	}
	var response models.Response
	response.FolloweeID = userID
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	err = database.RespondToFollow(response)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		jsonError := models.Error{
			Reason: err.Error(),
		}
		json.NewEncoder(w).Encode(jsonError)
		return
	}
	requestIDjson := struct {
		ID int `json:"id"`
	}{
		ID: response.ID,
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(requestIDjson)
}
