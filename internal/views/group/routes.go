package group

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupGroupRoutes() {
	http.HandleFunc("POST /api/group", middleware.AllowCORS(middleware.CheckAuth(CreateGroupHandler)))
	http.HandleFunc("GET /api/group/{id}", middleware.AllowCORS(middleware.CheckAuth(GetGroupPageHandler)))
	
	http.HandleFunc("POST /api/invitation", middleware.AllowCORS(middleware.CheckAuth(CreateInvitationHandler)))
	http.HandleFunc("POST /api/invitationresponse", middleware.AllowCORS(middleware.CheckAuth(InvitationResponseHandler)))
	http.HandleFunc("POST /api/group_request", middleware.AllowCORS(middleware.CheckAuth(CreateRequestHandler)))
	http.HandleFunc("POST /api/group_response", middleware.AllowCORS(middleware.CheckAuth(RequestResponseHandler)))
	// http.HandleFunc("POST /api/exitgroup", middleware.AllowCORS(middleware.CheckAuth(ExitGroupHandler)))
	// http.HandleFunc("/search_group", m(SearchGroupHandler))
	// http.HandleFunc("/create_event", m(CreateEventHandler))
	// http.HandleFunc("/event_response", m(EventResponseHandler))
}
