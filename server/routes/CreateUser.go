package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"vacation-planner/models"
)

// 1. Check to ensure this route is being called only via POST.
// 2. Initialize a structure for the response data.
// 3. Handle data from within request body to be acted upon.
// 4. Check the database for any user that already has the email.
// 5. Create the user if there has not been a user with the given email.
// 6. Check for any errors while creating new user.
// 7. Write and package success response based if the user was created.
// 8. Write and package failed response if the user was not created.

func (h DBRouter) CreateUser(w http.ResponseWriter, r *http.Request) {

	// 1.
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2.
	type SignupAttempt struct {
		Success bool   `json: "success"`
		Message string `json: "message"`
	}

	// 3.
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4.
	var user models.User
	result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

	// 5.
	if result.RowsAffected == 0 {
		user.Email = strings.ToLower(requestBody["Email"].(string))
		user.Password = requestBody["Password"].(string)
		if newUser := h.DB.Create(&user); newUser.Error != nil {
			// 6.
			fmt.Println(newUser.Error)
		}

		// 7.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := SignupAttempt{Success: true, Message: "User succesfully created account"}
		jsonResponse, err1 := json.Marshal(response)
		if err1 != nil {
			http.Error(w, err1.Error(), http.StatusBadRequest)
			return
		}
		w.Write(jsonResponse)

	} else {
		// 8.
		response := SignupAttempt{Success: false, Message: "Email is already in use"}
		jsonResponse, err2 := json.Marshal(response)
		if err2 != nil {
			http.Error(w, err2.Error(), http.StatusBadRequest)
			return
		}
		w.Write(jsonResponse)
	}
}
