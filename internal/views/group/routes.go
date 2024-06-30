package group

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupGroupRoutes() {
	http.HandleFunc("POST /api/create_group", middleware.ValidateSessionMiddleware(CreateGroupHandler))
	http.HandleFunc("GET /api/group/{id}", middleware.ValidateSessionMiddleware(GetGroupPageHandler))
	// http.HandleFunc("/invite_user", m(InvitationHandler))
	// http.HandleFunc("/group_request", m(RequestHandler))
	// http.HandleFunc("/search_group", m(SearchGroupHandler))
	// http.HandleFunc("/create_event", m(CreateEventHandler))
	// http.HandleFunc("/event_response", m(EventResponseHandler))
}
