package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"vacation-planner/database"
	"vacation-planner/routes"

	"github.com/gorilla/mux"
)

// Test function for GetDestInfo route
func TestGetDestInfo(t *testing.T) {
	db, err := database.Connect()
	h := routes.NewConnection(db)

	// Define the request
	req, err := http.NewRequest("GET", "/newDestination/Paris", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Define the response recorder
	rr := httptest.NewRecorder()

	// Create a new router and register the handler
	r := mux.NewRouter()
	r.HandleFunc("/newDestination/{location}", h.GetDestInfo).Methods("GET")

	// Serve the request and record the response
	r.ServeHTTP(rr, req)

	// Check the status code for 400 or 500 status
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
