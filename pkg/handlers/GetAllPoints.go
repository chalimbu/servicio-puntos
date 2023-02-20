package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) GetAllPoints(w http.ResponseWriter, r *http.Request) {
	var users []models.Point

	if result := h.DB.Find(&users); result.Error != nil {
		fmt.Println(result.Error)
	}

	pointsMap := make(map[int]int)
	for i := 0; i < len(users); i += 1 {
		pointsMap[users[i].Id] = users[i].Points
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pointsMap)
}
