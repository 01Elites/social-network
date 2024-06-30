package post

import ("net/http"
"social-network/internal/views/middleware")
func SetupPostRoutes(){
	http.HandleFunc("GET /api/posts", middleware.ValidateSessionMiddleware(GetPostsHandler))
	http.HandleFunc("GET /api/post/{id}", middleware.ValidateSessionMiddleware(GetPostByIDHandler))
	http.HandleFunc("GET /api/post/{id}/comments", middleware.ValidateSessionMiddleware(GetPostCommentsHandler))
	http.HandleFunc("POST /api/create_post", middleware.ValidateSessionMiddleware(CreatePostHandler))
	http.HandleFunc("POST /api/post/{id}/comments", middleware.ValidateSessionMiddleware(CreateCommentHandler))
	http.HandleFunc("POST /api/create_like/{id}", middleware.ValidateSessionMiddleware(CreateLikeHandler))
	http.HandleFunc("DELETE /api/post/{id}", middleware.ValidateSessionMiddleware(DeletePostHandler))
}