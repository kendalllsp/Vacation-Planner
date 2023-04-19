package routes

import (
	"encoding/json"
	"net/http"
	"vacation-planner/models"
)

// 1. Check to ensure this route is being called only via POST.
// 2. Initialize a structure for the response data.
// 3. Handle data from within request body to be acted upon.
// 4. Check the database to make sure any user has the email.
// 5. If there is a user with the email, check to see if the password given does not match that email's password.
// 6. Create, package and write failed response if the password does not match already created user with given email.
// 7. Create, package, and write success response if given password matches the saved password.
// 8. Create, package, and write failed response if no user in database has the given email.

func (h DBRouter) LoginUser(w http.ResponseWriter, r *http.Request) {

	// 1.
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// 2.
	type LoginAttempt struct {
		LoggedIn bool   `json: "loggedin"`
		Email    string `json: "email"`
		Message  string `json: "message"`
	}

	// 3.
	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 4.
	existingUser := &models.User{}
	result := h.DB.First(existingUser, "Email = ?", requestBody["Email"].(string))
	if result.RowsAffected != 0 {

		// 5.
		if existingUser.Password != requestBody["Password"].(string) {
    
			// 6.
			response := LoginAttempt{LoggedIn: false, Email: requestBody["Email"].(string), Message: "Email and password combination does not exist."}
			jsonResponse, err1 := json.Marshal(response)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
				return
			}
      w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)
      
    } else {
    
			// 7.
			response := LoginAttempt{LoggedIn: true, Email: requestBody["Email"].(string), Message: "User successfully logged in."}
			jsonResponse, err2 := json.Marshal(response)
			if err2 != nil {
				http.Error(w, err2.Error(), http.StatusBadRequest)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(jsonResponse)

		}
	} else {
  
		// 8.
		response := LoginAttempt{LoggedIn: false, Email: requestBody["Email"].(string), Message: "Email not in use in our userbase."}
		jsonResponse, err3 := json.Marshal(response)
		if err3 != nil {
			http.Error(w, err3.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonResponse)
	}
}
