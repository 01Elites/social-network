package pic

import (
	"net/http"
)

// servesImage handles gets the path of the image from url and serves it in the front end  as []byte
func servesImage(w http.ResponseWriter, r *http.Request) {
	//get the last part of the url
	path := r.URL.Path[len("/api/image/"):]
	path = "internal/database/images/" + path
	http.ServeFile(w, r, path)
}
