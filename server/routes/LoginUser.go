package routes

import (
	"encoding/json"
	"net/http"
	"vacation-planner/models"
)

// Login user POST, using HTTP request body information for email and password
func (h DBRouter) LoginUser(w http.ResponseWriter, r *http.Request) {

	// Only POST is allowed for this route
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Adding template response to be sent front end
	type LoginAttempt struct {
		LoggedIn bool   `json: "loggedin"`
		Email    string `json: "email"`
		Message  string `json: "message"`
	}

	// Creating new variable for storing request body
	var requestBody map[string]interface{}

	// Decoding body of the http request for the information for the user account and error checking
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Creating new variable to be referenced for current DB table, and new insertion to table
	existingUser := &models.User{}
	result := h.DB.First(existingUser, "Email = ?", requestBody["Email"].(string))

	// Checking if the rows that have the email isn't 0 therefore somebody has the email
	if result.RowsAffected != 0 {

		// Checking if the password given in the request does not match the password stored for the user in the DB
		if existingUser.Password != requestBody["Password"].(string) {
			// Creating new response under LoginAttempt struct style
			// Marshaling response as JSON and writing it as response
			// Now including the Email string in the response
			response := LoginAttempt{LoggedIn: false, Email: requestBody["Email"].(string), Message: "Email and password combination does not exist."}
			jsonResponse, err1 := json.Marshal(response)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
				return
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Returning the JSON response
			w.Write(jsonResponse)

		} else {
			// The password given in the requestBody does match the password stored for the email given.

			// Setting headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Initalizing a response based on the before mentioned response template with true loggedIn value and string details
			response := LoginAttempt{LoggedIn: true, Email: requestBody["Email"].(string), Message: "User successfully logged in."}

			// Packaging the response into a JSON and error checking
			jsonResponse, err2 := json.Marshal(response)
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusBadRequest)
				return
			}
			// Returning the JSON response
			w.Write(jsonResponse)

		}
	} else {
		// If Rows Affected (rows with email given) is  0, therefore nobody has an account with
		// the email given, we can't login the user and tell them their email is not in use by our website.

		response := LoginAttempt{LoggedIn: false, Email: requestBody["Email"].(string), Message: "Email not in use in our userbase."}

		// Packaging the response into a JSON
		jsonResponse, err3 := json.Marshal(response)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusBadRequest)
			return
		}

		// Returning the JSON
		w.Write(jsonResponse)
	}
}
