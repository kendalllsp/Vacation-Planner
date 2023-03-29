package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"vacation-planner/database"
	"vacation-planner/routes"

	"github.com/gorilla/mux"
)

func TestUpdateDestination(t *testing.T) {

	//format of returned JSON body
	type responseBody struct {
		Saved   bool   `json: "saved"`
		Message string `json: "message"`
	}

	//POST
	//case 1: no user with that email
	//set email that isnt in database
	email := "a@gmail.com"
	location := "Paris"
	payload := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email, location)

	//make a request to the database
	req, err := http.NewRequest("POST", "/updateDestination", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr := httptest.NewRecorder()

	db, err := database.Connect()
	h := routes.NewConnection(db)

	r := mux.NewRouter()
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("POST")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//the expected response is that the email is not in the database
	response := responseBody{Saved: false, Message: "No user with given email."}
	jsonResponse, err := json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("POST Case 1 passed!!\n")
	}

	//case 2: save new location (deleted in case 2 of DELETE)
	email2 := "123@gmail.com"
	location2 := "Oslo"
	payload2 := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email2, location2)
	req2, err := http.NewRequest("POST", "/updateDestination", strings.NewReader(payload2))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr2 := httptest.NewRecorder()
	r2 := mux.NewRouter()
	r2.HandleFunc("/updateDestination", h.UpdateDestination).Methods("POST")
	r2.ServeHTTP(rr2, req2)
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = responseBody{Saved: true, Message: "New location successfully saved."}
	jsonResponse, err = json.Marshal(response)
	if !bytes.Equal(rr2.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr2.Body.String(), jsonResponse)
	} else {
		fmt.Printf("POST Case 2 passed!!\n")
	}

	//case 3: location has already been saved
	email3 := "Test@test.com"
	location3 := "Gainesville, FL"
	payload3 := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email3, location3)
	req3, err := http.NewRequest("POST", "/updateDestination", strings.NewReader(payload3))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr3 := httptest.NewRecorder()
	r3 := mux.NewRouter()
	r3.HandleFunc("/updateDestination", h.UpdateDestination).Methods("POST")
	r3.ServeHTTP(rr3, req3)
	if status := rr3.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = responseBody{Saved: false, Message: "User already has location saved."}
	jsonResponse, err = json.Marshal(response)
	if !bytes.Equal(rr3.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr3.Body.String(), jsonResponse)
	} else {
		fmt.Printf("POST Case 3 passed!!\n")
	}

	//DELETE (just written bytes as of sprint 3)
	//case 1: no user with that email
	email4 := "a@gmail.com"
	location4 := "Paris"
	payload4 := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email4, location4)

	//make a request to the database
	req4, err := http.NewRequest("DELETE", "/updateDestination", strings.NewReader(payload4))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr4 := httptest.NewRecorder()
	r4 := mux.NewRouter()
	r4.HandleFunc("/updateDestination", h.UpdateDestination).Methods("DELETE")
	r4.ServeHTTP(rr4, req4)

	//error handling for wrong http status code
	if status := rr4.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//expected response:
	response = responseBody{Saved: false, Message: "No user with given email."}
	jsonResponse, err = json.Marshal(response)

	//testing that expected response matches rr4
	if !bytes.Equal(rr4.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr4.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 1 passed!!\n")
	}

	//case 2: location successfully deleted
	email5 := "123@gmail.com"
	location5 := "Oslo"
	payload5 := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email5, location5)
	req5, err := http.NewRequest("DELETE", "/updateDestination", strings.NewReader(payload5))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr5 := httptest.NewRecorder()
	r5 := mux.NewRouter()
	r5.HandleFunc("/updateDestination", h.UpdateDestination).Methods("DELETE")
	r5.ServeHTTP(rr5, req5)
	if status := rr5.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = responseBody{Saved: false, Message: "Location successfuly deleted."}
	jsonResponse, err = json.Marshal(response)
	if !bytes.Equal(rr5.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr5.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 2 passed!!\n")
	}

	//case 3: account does not have location saved
	email6 := "123@gmail.com"
	location6 := "Helsinki"
	payload6 := fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email6, location6)
	req6, err := http.NewRequest("DELETE", "/updateDestination", strings.NewReader(payload6))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr6 := httptest.NewRecorder()
	r6 := mux.NewRouter()
	r6.HandleFunc("/updateDestination", h.UpdateDestination).Methods("DELETE")
	r6.ServeHTTP(rr6, req6)
	if status := rr6.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	response = responseBody{Saved: false, Message: "No saved location by the specfied user matches the location given."}
	jsonResponse, err = json.Marshal(response)
	if !bytes.Equal(rr6.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr6.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 3 passed!!\n")
	}

	//GET (just written bytes as of sprint 3)
	//case 1: no user with that email
	email = "a@gmail.com"
	location = "Paris"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email, location)

	//make a request to the database
	req, err = http.NewRequest("GET", "/updateDestination", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("GET")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//expected response:
	expectedResponse := []byte("No user with the email address associated.")

	//testing that expected response matches rr4
	if !bytes.Equal(rr.Body.Bytes(), expectedResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), expectedResponse)
	} else {
		fmt.Printf("GET Case 1 passed!!\n")
	}

	//case 2: returns an array of locations
	email = "Test@test.com"
	location = "Paris"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email, location)
	req, err = http.NewRequest("GET", "/updateDestination", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		fmt.Printf("GET Case 2 passed!!\n")
	}

	/**expectedResponse = []byte("Gainesville, Miami")
	//expected response is an array of locations
	var locations []models.SavedLocation := ({Email: })

	if !bytes.Equal(rr.Body.Bytes(), expectedResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), expectedResponse)
	} else {
		fmt.Printf("GET Case 2 passed!!\n")
	}**/

	//case 3: list empty
	email = "123@gmail.com"
	location = "Paris"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s"}`, email, location)
	req, err = http.NewRequest("GET", "/updateDestination", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		fmt.Printf("GET Case 3 passed!!\n")
	}

	/**expectedResponse = []byte("Gainesville, Miami")
	//expected response is an array of locations
	var locations []models.SavedLocation := ({Email: })

	if !bytes.Equal(rr.Body.Bytes(), expectedResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), expectedResponse)
	} else {
		fmt.Printf("GET Case 2 passed!!\n")
	}**/

}
