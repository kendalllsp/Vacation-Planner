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

func TestCreateUser(t *testing.T) {
	//test against validation strings
	//testing strings may have to be slightly changed for each test of create user

	//establishing response struct
	type SignupAttempt struct {
		Success bool   `json: "success"`
		Message string `json: "message"`
	}

	//case 1: email taken
	//this email is in the database already
	email := "123@gmail.com"
	password := "123"

	payload := fmt.Sprintf(`{"Email": "%s", "Password": "%s"}`, email, password)

	//making a request to the database
	req, err := http.NewRequest("POST", "/createUser", strings.NewReader(payload))
	if err != nil {
		t.Errorf("Error: create user request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr := httptest.NewRecorder()

	db, err := database.Connect()
	h := routes.NewConnection(db)

	r := mux.NewRouter()
	r.HandleFunc("/createUser", h.CreateUser).Methods("POST")
	r.ServeHTTP(rr, req)

	//error handling for wrong http status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//with this email and password, we expect the user to be logged in
	response := SignupAttempt{Success: false, Message: "Email is already in use"}
	jsonResponse, err := json.Marshal(response)

	//testing if the response matches output
	if !bytes.Equal(rr.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr.Body.String(), jsonResponse)
	} else {
		fmt.Printf("Case 1 passed!!\n")
	}

	//case 2: user created successfully

	//email must be changed to one not in database for this to work
	email2 := "klspab@gmail.com"
	password2 := "12345"

	payload2 := fmt.Sprintf(`{"Email": "%s", "Password": "%s"}`, email2, password2)

	//making a request to the database
	req2, err := http.NewRequest("POST", "/createUser", strings.NewReader(payload2))
	if err != nil {
		t.Errorf("Error: create user request could not be completed")
	}

	//response recorder: takes in the returned bytestring
	rr2 := httptest.NewRecorder()

	r2 := mux.NewRouter()
	r2.HandleFunc("/createUser", h.CreateUser).Methods("POST")
	r2.ServeHTTP(rr2, req2)

	//error handling for wrong http status code
	if status := rr2.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	//with this email and password, we expect the user to be logged in
	response2 := SignupAttempt{Success: true, Message: "User succesfully created account"}
	jsonResponse, err = json.Marshal(response2)

	//testing if the response matches output
	if !bytes.Equal(rr2.Body.Bytes(), jsonResponse) {
		t.Errorf("handler returned unexpected response: got %v want %v", rr2.Body.String(), jsonResponse)
	} else {
		fmt.Printf("Case 2 passed!!\n")
	}
}
