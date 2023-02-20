package handlers

import (
	"encoding/json"
	"net/http"
)

func (h handler) Health(w http.ResponseWriter, r *http.Request) {

	pointsMap := make(map[string]string)
	pointsMap["status"] = "ok"

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pointsMap)
}
