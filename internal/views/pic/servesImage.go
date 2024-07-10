package pic

import (
	"net/http"
	"path/filepath"
)

// servesImage handles gets the path of the image from url and serves it in the front end  as []byte
func servesImage(w http.ResponseWriter, r *http.Request) {
	//get the last part of the url
	fileName := r.URL.Path[len("/api/image/"):]
	path := filepath.Join("internal/database/images", fileName)
	http.ServeFile(w, r, path)
}
