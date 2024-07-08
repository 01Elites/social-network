package group

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupGroupRoutes() {
	http.HandleFunc("POST /api/create_group", middleware.AllowCORS(middleware.CheckAuth(CreateGroupHandler)))
	http.HandleFunc("GET /api/group/{id}", middleware.AllowCORS(middleware.CheckAuth(GetGroupPageHandler)))

	http.HandleFunc("POST /api/invitation", middleware.AllowCORS(middleware.CheckAuth(CreateInvitationHandler)))
	http.HandleFunc("PATCH /api/invitation_response", middleware.AllowCORS(middleware.CheckAuth(InvitationResponseHandler)))

	http.HandleFunc("POST /api/join_group_req", middleware.AllowCORS(middleware.CheckAuth(CreateRequestHandler)))
	http.HandleFunc("PATCH /api/cancel_join_req", middleware.AllowCORS(middleware.CheckAuth(CancelRequestHandler)))
	http.HandleFunc("PATCH /api/join_group_res", middleware.AllowCORS(middleware.CheckAuth(RequestResponseHandler)))

	http.HandleFunc("POST /api/exit_group", middleware.AllowCORS(middleware.CheckAuth(ExitGroupHandler)))

	http.HandleFunc("POST /api/create_event", middleware.AllowCORS(middleware.CheckAuth(CreateEventHandler)))
	// http.HandleFunc("GET /api/event/{id}", middleware.AllowCORS(middleware.CheckAuth(GetEventHandler)))
	http.HandleFunc("POST /api/event_response", middleware.AllowCORS(middleware.CheckAuth(EventResponseHandler)))

}
