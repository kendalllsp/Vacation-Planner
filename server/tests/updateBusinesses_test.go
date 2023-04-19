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

func TestUpdateBusinesses(t *testing.T) {

	//format of returned JSON body
	type responseBody struct {
		Saved   bool   `json: "saved"`
		Message string `json: "message"`
	}

	//POST
	//case 1: no user with given email
	email := "a@gmail.com"
	location := "Austin, TX"
	name := "BBQ Deez Nuts"
	price := "$"
	rating := 5.5
	blocation := "123 BLM Dr."
	placetype := "BBQ"

	payload := fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)

	//make a request to the database
	req, err := http.NewRequest("POST", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr := httptest.NewRecorder()

	db, err := database.Connect()
	h := routes.NewConnection(db)

	r := mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("POST")
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

	//case 2: new business successfully saved
	email = "123@gmail.com"
	location = "Austin, TX"
	name = "Kendy's BBQ"
	price = "$"
	rating = 5.5
	blocation = "1234 kendy street"
	placetype = "BBQ"

	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	//make a request to the database
	req, err = http.NewRequest("POST", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("POST")
	r.ServeHTTP(rr, req)
	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//the expected response is that the email is not in the database
	response = responseBody{Saved: true, Message: "New business successfully saved."}
	jsonResponse, err = json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("POST Case 2 passed!!\n")
	}

	//case 3: location already saved
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	//make a request to the database
	req, err = http.NewRequest("POST", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("POST")
	r.ServeHTTP(rr, req)
	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//the expected response is that the email is not in the database
	response = responseBody{Saved: false, Message: "User already has location saved."}
	jsonResponse, err = json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("POST Case 3 passed!!\n")
	}

	//DELETE
	//case 1: no user with email
	email = "a@gmail.com"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)

	//make a request to the database
	req, err = http.NewRequest("DELETE", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("DELETE")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//the expected response is that the email is not in the database
	response = responseBody{Saved: false, Message: "No user with given email."}
	jsonResponse, err = json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 1 passed!!\n")
	}

	//case 2: location successfully deleted (delete one saved in post case 2)
	email = "123@gmail.com"
	location = "Austin, TX"
	name = "Kendy's BBQ"
	price = "$"
	rating = 5.5
	blocation = "1234 kendy street"
	placetype = "BBQ"

	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	//make a request to the database
	req, err = http.NewRequest("DELETE", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("DELETE")
	r.ServeHTTP(rr, req)
	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//the expected response is that the email is not in the database
	response = responseBody{Saved: false, Message: "Location successfuly deleted."}
	jsonResponse, err = json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 2 passed!!\n")
	}

	//case 3: no such location
	email = "123@gmail.com"
	location = "Austin, TX"
	name = "Not Real"
	price = "$"
	rating = 5.5
	blocation = "1234 kendy street"
	placetype = "BBQ"

	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	//make a request to the database
	req, err = http.NewRequest("DELETE", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("DELETE")
	r.ServeHTTP(rr, req)
	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
	//the expected response is that the email is not in the database
	response = responseBody{Saved: false, Message: "No saved location by the specfied user matches the location given."}
	jsonResponse, err = json.Marshal(response)

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("DELETE Case 3 passed!!\n")
	}

	//GET
	//case 1: no user with email
	email = "a@gmail.com"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)

	//make a request to the database
	req, err = http.NewRequest("GET", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("GET")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//the expected response is that the email is not in the database
	expectedResponse := []byte("No user with the email address associated.")

	//testing the the response matches expected output
	if !bytes.Equal(rr.Body.Bytes(), expectedResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), expectedResponse)
	} else {
		fmt.Printf("GET Case 1 passed!!\n")
	}

	//case 2: string of saved locations
	email = "123@gmail.com"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	req, err = http.NewRequest("GET", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		fmt.Printf("GET Case 2 passed!!\n")
	}

	//case 3: empty string array
	email = "123@gmail.com"
	location = "Tampa, FL"
	payload = fmt.Sprintf(`{"Email": "%s", "Location": "%s", "Name": "%s", "Price": "%s", "Rating": %f, "B_Location": "%s", "Type": "%s"}`, email, location, name, price, rating, blocation, placetype)
	req, err = http.NewRequest("GET", "/updateBusinessList", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: request could not be completed")
	}
	rr = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("GET")
	r.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	} else {
		fmt.Printf("GET Case 3 passed!!\n")
	}
}
