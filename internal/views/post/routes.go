package post

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupPostRoutes() {

	/****************************************** GET Requests *********************************************/
	http.HandleFunc("GET /api/posts", middleware.AllowCORS(middleware.CheckAuth(GetPostsHandler)))
	http.HandleFunc("GET /api/post/{id}", middleware.AllowCORS(middleware.CheckAuth(GetPostByIDHandler)))
	http.HandleFunc("GET /api/post/{id}/comments", middleware.AllowCORS(middleware.CheckAuth(GetPostCommentsHandler)))

	/****************************************** POST Requests *********************************************/
	http.HandleFunc("POST /api/post", middleware.AllowCORS(middleware.CheckAuth(CreatePostHandler)))
	http.HandleFunc("POST /api/post/{id}/comments", middleware.AllowCORS(middleware.CheckAuth(CreateCommentHandler)))
	http.HandleFunc("POST /api/like", middleware.AllowCORS(middleware.CheckAuth(CreateLikeHandler)))

	/****************************************** DELETE Requests *********************************************/
	http.HandleFunc("DELETE /api/post/{id}", middleware.AllowCORS(middleware.CheckAuth(DeletePostHandler)))
}
