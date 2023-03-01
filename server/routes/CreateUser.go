package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vacation-planner/models"
)

// Create user POST, using HTTP request body information for email and password
func (h DBRouter) CreateUser(w http.ResponseWriter, r *http.Request) {

	// Only POST is allowed for this route
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Adding template response for front end
	type SignupAttempt struct {
		Success 	bool   `json: "success"`
		Message 	string `json: "message"`
	}

	// Creating two new variables to use as reference
	var requestBody map[string]interface{}
	var user models.User

	// Decoding body of the http request for the information for the user account
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Added a line to check the database for any users with the same email as the new account

	result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

	// Checking if the rows that have the email is 0 therefore nobody has the email
	if result.RowsAffected == 0 {
		// Assigning Email and Password to new User
		user.Email = requestBody["Email"].(string)
		user.Password = requestBody["Password"].(string)

		// Creating new user in the DB and checking for error
		if newUser := h.DB.Create(&user); newUser.Error != nil {
			fmt.Println(newUser.Error)
		}

		// Setting headers
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Initalizing response variable to hold a boolean and a string message
		response := SignupAttempt{ Success: true, Message: "User succesfully created account" }
		// Packing response as type JSON
		jsonResponse, err1 := json.Marshal(response)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		// Write JSON response
		w.Write(jsonResponse)

	} else {
		// If Rows Affected (rows with email given) is greater than 0, therefore someone has an account with
		// the email given, we don't create a new user and tell them their email is taken.
		response := SignupAttempt{ Success: false, Message: "Email is already in use" }
		jsonResponse, err2 := json.Marshal(response)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		w.Write(jsonResponse)
	}
}
