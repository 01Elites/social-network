package post

import (
	"net/http"

	"social-network/internal/views/middleware"
)

func SetupPostRoutes() {
	http.HandleFunc("GET /api/posts", middleware.AllowCORS(middleware.CheckAuth(GetPostsHandler)))
	http.HandleFunc("GET /api/post/{id}", middleware.AllowCORS(middleware.CheckAuth(GetPostByIDHandler)))
	http.HandleFunc("GET /api/post/{id}/comments", middleware.AllowCORS(middleware.CheckAuth(GetPostCommentsHandler)))
	http.HandleFunc("POST /api/post", middleware.AllowCORS(middleware.CheckAuth(CreatePostHandler)))
	http.HandleFunc("POST /api/post/{id}/comments", middleware.AllowCORS(middleware.CheckAuth(CreateCommentHandler)))
	http.HandleFunc("POST /api/like/{id}", middleware.AllowCORS(middleware.CheckAuth(CreateLikeHandler)))
	http.HandleFunc("DELETE /api/post/{id}", middleware.AllowCORS(middleware.CheckAuth(DeletePostHandler)))
}
