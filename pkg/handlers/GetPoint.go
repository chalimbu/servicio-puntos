package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) GetPoint(w http.ResponseWriter, r *http.Request) {
	// Read dynamic id parameter
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	// Find User points
	var user models.Point

	if result := h.DB.First(&user, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	pointsMap := make(map[int]int)
	pointsMap[user.IdUser] = user.Points

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(pointsMap)
}
