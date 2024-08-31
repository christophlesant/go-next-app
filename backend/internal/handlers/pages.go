package handlers

import (
	"encoding/json"
	"net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "/go.svg"} // Show logo if front end is connected!
	json.NewEncoder(w).Encode(response)
}
