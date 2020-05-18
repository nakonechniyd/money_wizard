package apiv1

import (
	"encoding/json"
	"net/http"
)

// ErrorResponse -
type ErrorResponse struct {
	Error string
}

// ErrorJSON -
func ErrorJSON(w http.ResponseWriter, msg string) {
	json.NewEncoder(w).Encode(ErrorResponse{Error: msg})
}
