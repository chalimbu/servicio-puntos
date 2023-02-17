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

	var user models.User
	json.Unmarshal(body, &user)

	log.Println(user)
	//get the actual points of the user
	var currentUser models.User

	if result := h.DB.First(&currentUser, user.Id); result.Error != nil {
		log.Println(result.Error)
	}

	log.Println("current user")
	log.Println(currentUser)

	// Append to the Books table
	//if result := h.DB.Create(&book); result.Error != nil {
	//	fmt.Println(result.Error)
	//}

	// Send a 201 created response
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")
}
