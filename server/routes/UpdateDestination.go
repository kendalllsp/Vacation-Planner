package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"vacation-planner/models"
)

// Update Destination route, using HTTP method (Get/Post/Delete) to determine what
// action to do with body information updating saved locations
func (h DBRouter) UpdateDestination(w http.ResponseWriter, r *http.Request) {

	// Creating two new variables
	// requestBody to store the body of the HTTP request and refer to it
	// savedLocation to add to the database of savedLocations
	var requestBody map[string]interface{}

	// Initializing Struct for JSON response data
	type responseBody struct {
		Saved 	bool 	`json: "saved"`
		Message string 	`json: "message"`
	}

	// If adding to location list
	if r.Method == "POST" {

		var savedLocation models.SavedLocation

		// Decoding body of the http request for the information for the newly saved location
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checking the database for a user with the same email as the account trying to update location list
		result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

		// If there is no user with said email, return error
		if result.RowsAffected == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			
			// Creating new response body based on the situation and returning it
			response := responseBody { Saved: false, Message: "No user with given email." }

			jsonResponse, err1 := json.Marshal(response)
			if err1 != nil {
				http.Error(w, err1.Error(), http.StatusBadRequest)
				return
			}
			w.Write(jsonResponse)

		} else {

			// Checking the savedLocations for a value with the email and location, meaning the wrong call was called.. meaning to delete already saved location
			result = h.DB.Where(&models.SavedLocation{Email: requestBody["Email"].(string), Location: requestBody["Location"].(string)}).First(&models.SavedLocation{})

			// Checking if the rows that have the email is 0 therefore they have not already saved given location
			if result.RowsAffected == 0 {

				// Assigning Email and Location to new location
				savedLocation.Email = requestBody["Email"].(string)
				savedLocation.Location = requestBody["Location"].(string)
	
				// Creating new location in the DB and checking for error
				if newLocation := h.DB.Create(&savedLocation); newLocation.Error != nil {
					fmt.Println(newLocation.Error)
				}

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it

				response := responseBody { Saved: true, Message: "New location successfully saved." }

				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
				w.Write(jsonResponse)

			} else {
				// If Rows Affected (rows with email given) is greater than 0, therefore the user with said email
				// has already saved the location given. In order to delete, front end should call Delete HTTP request not PUT.
				
				// Creating new response body based on the situation and returning it

				response := responseBody { Saved: false, Message: "User already has location saved." }
				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
				w.Write(jsonResponse)
			}
		}
	// if deleting from location list
	} else if r.Method == "DELETE" {

		// Decoding body of the http request for the information of the wanted to be deleted location
		err := json.NewDecoder(r.Body).Decode(&requestBody)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Checking user table for user with email given
		result := h.DB.First(&models.User{}, "Email = ?", requestBody["Email"].(string))

		// If no user has given email
		if result.RowsAffected == 0 {
			w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it

				response := responseBody { Saved: false, Message: "No user with given email." }

				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
				w.Write(jsonResponse)
		} else {
			// Checking the savedLocations for a value with the email and location
			result = h.DB.Where(&models.SavedLocation{Email: requestBody["Email"].(string), Location: requestBody["Location"].(string)}).First(&models.SavedLocation{})

			// Checking if the rows that have the email/location is 0 therefore they have not already saved given location
			if result.RowsAffected != 0 {

				if deleteLocation := result.Delete(&models.SavedLocation{}); deleteLocation.Error != nil {
					fmt.Println(deleteLocation.Error)
				}

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it

				response := responseBody { Saved: false, Message: "Location successfuly deleted." }

				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
				w.Write(jsonResponse)

			} else {
				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)

				// Creating new response body based on the situation and returning it

				response := responseBody { Saved: false, Message: "No saved location by the specfied user matches the location given." }

				jsonResponse, err1 := json.Marshal(response)
				if err1 != nil {
					http.Error(w, err1.Error(), http.StatusBadRequest)
					return
				}
				w.Write(jsonResponse)
			}
		}
	} else if r.Method == "GET" {

		// Changed GET request to refer to url parameters
		// and changed the rest of the GET's email references to this variable
		email := r.URL.Query().Get("Email")

		// Checking the database for a user with the same email as the account trying to update location list
		result := h.DB.First(&models.User{}, "Email = ?", email)

		// If there is no user with said email, return error
		if result.RowsAffected == 0 {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("No user with the email address associated."))
		} else {
			// Start a new variable that is a slice of saved locations, to add the specific users locations to
			var locations []models.SavedLocation
			// Finding all rows within Saved Locations with the given email
			h.DB.Where("email = ?", email).Find(&locations)

			// Checking if the rows that have the email is not 0 therefore they have saved locations
			if len(locations) != 0 {

				// Setting headers
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
		
				// Send the location slice to the front end
				json.NewEncoder(w).Encode(locations)

			} else {
				// If Rows Affected (rows with email given) is 0, the user has no saved locations.
				w.Write([]byte("User destination list is empty."))
			}
		}
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
