package main

import (
	"fmt"
	"log"
	"net/http"
	"vacation-planner/database"
	"vacation-planner/routes"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Establishing database connection
	db, err := database.Connect()
	if err != nil {
		fmt.Println("Error: could not connect to database")
	}

	// Packaging database connection to be passed to handlers
	h := routes.NewConnection(db)

	r := mux.NewRouter()

	// Route that can take 3 HTTP method requests to handle saved businesses in the savedBusinesses table for users to have businesses saved on their own Destination Lists
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("POST")
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Methods("DELETE")
	r.HandleFunc("/updateBusinessList", h.UpdateBusinessList).Queries("email", "{email}", "location", "{location}").Methods("GET")

	// POST request for creating a user in the database
	r.HandleFunc("/createUser", h.CreateUser).Methods("POST")

	// GET request to get relevant travel information for the user
	r.HandleFunc("/newDestination", h.GetDestInfo).Queries("location", "{location}", "start", "{start}", "end", "{end}").Methods("GET")

	// POST to login a user while using username and password. Post ensures the info is not on the URL
	r.HandleFunc("/loginUser", h.LoginUser).Methods("POST")

	// New route that can take 3 HTTP method requests to handle saved locations in the savedLocation table for users to have their own Destination Lists
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("POST")
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("DELETE")
	r.HandleFunc("/updateDestination", h.UpdateDestination).Methods("GET")

	// Enabling CORS, binding to a port and passing our router in
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"http://localhost:4200"})
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
