package helpers

import (
    "encoding/json"
    "net/http"
)

// HTTPErrorResponse represents a structured error response for HTTP requests.
type HTTPErrorResponse struct {
    Reason string `json:"reason"`
}

// SendHTTPError sends a JSON formatted HTTP error response with a specified status code.
func HTTPError(w http.ResponseWriter, message string, statusCode int) {
    w.WriteHeader(statusCode)
    errResponse := HTTPErrorResponse{Reason: message}
    json.NewEncoder(w).Encode(errResponse)
}
