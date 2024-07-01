package handlers

import (
	"encoding/json"
	"net/http"
)

type CheckResponse struct {
	Status string `json:"Status"`
}

func CheckHandler(w http.ResponseWriter, r *http.Request) {

	response := CheckResponse{Status: "OK"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(response)

}
