package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tutorials/go/crud/pkg/models"
)

func (h handler) ModifyPoints(w http.ResponseWriter, r *http.Request) {
	// Read to request body
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var user models.Point
	json.Unmarshal(body, &user)

	log.Println(user)
	//get the actual points of the user
	var currentUser models.Point

	if result := h.DB.First(&currentUser, user.IdUser); result.Error != nil {
		log.Println(result.Error)
	}

	log.Println("current user")
	log.Println(currentUser)

	newPoints := currentUser.Points + user.Points
	if newPoints >= 0 {
		currentUser.Points = newPoints

		h.DB.Save(&currentUser)

		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"status": "update"})
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"status": "failed"})
	}
}
