package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vacation-planner/models"
)

// Update Destination route, using HTTP method (Get/Post/Delete) to determine what
// action to do with request information when dealing with saved locations
func (h DBRouter) UpdateBusinessList(w http.ResponseWriter, r *http.Request) {

	// Creating a request body variable to store the request body information
	var requestBody map[string]interface{}

	// Initializing Struct template for JSON response data
	type responseBody struct {
		Saved   bool   `json: "saved"`
		Message string `json: "message"`
	}

	// If POST, then user is adding to their location list
	if r.Method == "POST" {

		// Creating a variable to be referenced by the database for existing table, and inserting into said table
		var savedBusiness models.SavedBusiness

		// Decoding body of the http request for the information for the newly saved location and error checking
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checking the database for the first user with the same email as the account trying to update location list
		result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

		// If there is no user with said email, return error
		if result.RowsAffected == 0 {

			// Setting headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Creating new response body based on the template struct, and the situation and returning it
			response := responseBody{Saved: false, Message: "No user with given email."}

			// Packaging to JSON
			jsonResponse, err1 := json.Marshal(response)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
				return
			}

			// Returning the JSON
			w.Write(jsonResponse)

		} else {

			// Checking the savedLocations for a value with the email and location, meaning the user has already saved the location
			result = h.DB.Where(&models.SavedBusiness{Email: requestBody["Email"].(string), Location: requestBody["Location"].(string), Name: requestBody["Name"].(string), B_Location: requestBody["B_Location"].(string)}).First(&models.SavedBusiness{})

			// Checking if the rows that have the email and location is 0 therefore they have not already saved given business
			if result.RowsAffected == 0 {

				// Assigning Email and Location to new location
				savedBusiness.Email = requestBody["Email"].(string)
				savedBusiness.Location = requestBody["Location"].(string)
				savedBusiness.Name = requestBody["Name"].(string)
				savedBusiness.Price = requestBody["Price"].(string)
				savedBusiness.Rating = requestBody["Rating"].(float64)
				savedBusiness.B_Location = requestBody["B_Location"].(string)
				savedBusiness.Type = requestBody["Type"].(string)

				// Creating new location in the DB and checking for error
				if newLocation := h.DB.Create(&savedBusiness); newLocation.Error != nil {
					fmt.Println(newLocation.Error)
				}

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it
				response := responseBody{Saved: true, Message: "New business successfully saved."}

				// Packaging into JSON
				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}

				// Returning JSON
				w.Write(jsonResponse)

			} else {
				// If Rows Affected (rows with email given) is not 0, user with said email has already saved the location given.

				// Creating new response body based on the situation and returning it
				response := responseBody{Saved: false, Message: "User already has location saved."}

				// Packaging to JSON
				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}

				// Returning JSON
				w.Write(jsonResponse)

			}
		}
		// If DELETE, the user is deleting from their location list
	} else if r.Method == "DELETE" {

		// Decoding body of the http request for the information of the wanted to be deleted location and error checking
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checking user table for user with email given
		result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

		// If rows = 0, then no user has the given email
		if result.RowsAffected == 0 {

			// Setting headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Creating new response body based on the situation and returning it
			response := responseBody{Saved: false, Message: "No user with given email."}

			// Packaging response to JSON
			jsonResponse, err1 := json.Marshal(response)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
				return
			}

			// Returning JSON response
			w.Write(jsonResponse)

		} else {
			// There is a user with the given email in the database

			// Checking the savedLocations for a value with the email and location
			result = h.DB.Where(&models.SavedBusiness{Email: requestBody["Email"].(string), Location: requestBody["Location"].(string), Name: requestBody["Name"].(string), B_Location: requestBody["B_Location"].(string)}).First(&models.SavedBusiness{})

			// If the rows that have the email/location is not 0 therefore they have already saved given location
			if result.RowsAffected != 0 {

				// Attempt to delete the user's saved locations in the database and error check
				if deleteLocation := result.Delete(&models.SavedBusiness{}); deleteLocation.Error != nil {
					fmt.Println(deleteLocation.Error)
				}

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it
				response := responseBody{Saved: false, Message: "Location successfuly deleted."}

				// Packaging to JSON
				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}

				// Returning JSON
				w.Write(jsonResponse)

			} else {
				// Else the rows that have the email/location is 0 therefore they have not already saved given location

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it
				response := responseBody{Saved: false, Message: "No saved location by the specfied user matches the location given."}

				// Packaging to JSON
				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}

				// Returning JSON
				w.Write(jsonResponse)
			}
		}
		// If GET, the user is attempting to react the "/trips" page and we return all saved locations for the user
	} else if r.Method == "GET" {

		// Accessing the user's email from the URL parameters (GET)
		email := r.URL.Query().Get("Email")

		// Accessing the user's email from the URL parameters (GET)
		location := r.URL.Query().Get("Location")

		// Checking the database for the first user with email trying to update their location list
		result := h.DB.First(&models.User{}, "Email = ?", email)

		// If there is no user with said email, return error
		if result.RowsAffected == 0 {

			// Setting headers
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// Returing byte slice with String details of results.
			w.Write([]byte("No user with the email address associated."))

		} else {

			// Initializing variable that is a slice of saved locations, to add the specific users locations to
			var businesses []models.SavedBusiness

			// Finding all rows within Saved Locations with the given email
			h.DB.Where("email = ?", email).Where("location = ?", location).Find(&businesses)

			// Checking if the rows that have the email is not 0 therefore they have saved locations
			if len(businesses) != 0 {

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Send the location slice to the front end
				json.NewEncoder(w).Encode(businesses)

			} else {
				// If Rows Affected (rows with email given) is 0, the user has no saved locations.

				// Setting header
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Initializing an empty string slice to avoid front end errors of an empty response
				var strings [0]string

				// Return empty string slice
				json.NewEncoder(w).Encode(strings)
			}
		}
		// If the method is neither GET, DELETE, or POST, then throw an error
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
