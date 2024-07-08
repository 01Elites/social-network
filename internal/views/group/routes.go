package group

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupGroupRoutes() {
	http.HandleFunc("POST /api/group", middleware.AllowCORS(middleware.CheckAuth(CreateGroupHandler)))
	http.HandleFunc("GET /api/group/{id}", middleware.AllowCORS(middleware.CheckAuth(GetGroupPageHandler)))

	http.HandleFunc("POST /api/invitation", middleware.AllowCORS(middleware.CheckAuth(CreateInvitationHandler)))
	http.HandleFunc("PATCH /api/invitationresponse", middleware.AllowCORS(middleware.CheckAuth(InvitationResponseHandler)))

	http.HandleFunc("POST /api/group_request", middleware.AllowCORS(middleware.CheckAuth(CreateRequestHandler))) //change naming to more standrad format
	// /api/join_group_req
	http.HandleFunc("PATCH /api/group_response", middleware.AllowCORS(middleware.CheckAuth(RequestResponseHandler))) //change naming to more standrad format
	// /api/join_group_res
	http.HandleFunc("PATCH /api/cancel_join_req", middleware.AllowCORS(middleware.CheckAuth(CancelRequestHandler)))
	http.HandleFunc("POST /api/exitgroup", middleware.AllowCORS(middleware.CheckAuth(LeaveGroupHandler)))

	http.HandleFunc("POST /api/create_event",middleware.AllowCORS(middleware.CheckAuth(CreateEventHandler)))

	// http.HandleFunc("/create_event", m(CreateEventHandler))
	// http.HandleFunc("/event_response", m(EventResponseHandler))


	// http.HandleFunc("POST /api/cancelinvite", middleware.AllowCORS(middleware.CheckAuth(CreateInvitationHandler))) "not needed"
	// http.HandleFunc("/search_group", m(SearchGroupHandler)) "not needed"
}
