package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"vacation-planner/models"

	"github.com/gorilla/mux"
	"github.com/jmatth11/yfusion"
	"github.com/joho/godotenv"
)

func (h DBRouter) GetDestInfo(w http.ResponseWriter, r *http.Request) {

	// Setting the Headers for the response type to be JSON
	w.Header().Set("Content-Type", "application/json")

	// Loading the API key from a private .env file with GoDotEnv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}

	// Setting the local API key value
	yelpAPIKey := os.Getenv("YELP_REST_API_KEY")

	// Taking the location value typed by the user, to use for computations
	params := mux.Vars(r)
	destinationLocation := params["location"]

	// Starting new Yelp client
	yelp := yfusion.NewYelpFusion(yelpAPIKey)

	// Setting all relevant filters for Fashion store search
	businessSearch := &yfusion.BusinessSearchParams{}
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("fashion")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// Searching based off of previous filters for Fashion search and error checking
	shoppingResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		fmt.Println("Fashion clothing business search could not be completed.")
	}

	// Creating new slice of size 10 for the top 10 rated Fashion Shopping businesses
	var ShoppingList [10]models.Business

	// Looping through results of the search to populate slice of fashion shopping businesses
	for i, b := range shoppingResult.Businesses {
		ShoppingList[i].Name = b.Name
		ShoppingList[i].Price = b.Price
		ShoppingList[i].Rating = b.Rating
		ShoppingList[i].Location = b.Location.Address1
		ShoppingList[i].Type = b.Categories[0].Title

		// Filling price string with not listed rather than empty string
		if len(ShoppingList[i].Price) < 1 {
			ShoppingList[i].Price = "Not listed"
		}
	}

	// Entertainment search filters
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("arts")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// Entertainment search results based on before mentioned filters for Entertainment
	entertainmentResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		log.Fatal(err)
	}

	// Entertainment slice
	var EntertainmentList [10]models.Business

	// Entertainment slice population
	for i, b := range entertainmentResult.Businesses {
		EntertainmentList[i].Name = b.Name
		EntertainmentList[i].Price = b.Price
		EntertainmentList[i].Rating = b.Rating
		EntertainmentList[i].Location = b.Location.Address1
		EntertainmentList[i].Type = b.Categories[0].Title

		// "Not listed" when the price is empty
		if len(EntertainmentList[i].Price) < 1 {
			EntertainmentList[i].Price = "Not listed"
		}
	}

	// Food search filters
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("food")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// Food search results based off of food search filters
	restaurantResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		log.Fatal(err)
	}

	// Food slice
	var RestaurantList [10]models.Business

	// Food slice populating
	for i, b := range restaurantResult.Businesses {
		RestaurantList[i].Name = b.Name
		RestaurantList[i].Price = b.Price
		RestaurantList[i].Rating = b.Rating
		RestaurantList[i].Location = b.Location.Address1
		RestaurantList[i].Type = b.Categories[0].Title

		// "Not listed" when the price is empty
		if len(RestaurantList[i].Price) < 1 {
			RestaurantList[i].Price = "Not listed"
		}
	}

	// Creating variables to calculate formal city, state, country for the response
	// Using the top rated business with the shortest distance from the designated location
	// To determine city, state and country
	var shortestDistance float64
	var city string
	var state string
	var country string

	for index := 0; index < len(restaurantResult.Businesses); index++ {
		if index == 0 {
			shortestDistance = restaurantResult.Businesses[index].Distance
			city = restaurantResult.Businesses[index].Location.City
			state = restaurantResult.Businesses[index].Location.State
			country = restaurantResult.Businesses[index].Location.Country
		} else {
			if shortestDistance >= restaurantResult.Businesses[index].Distance {
				shortestDistance = restaurantResult.Businesses[index].Distance
				city = restaurantResult.Businesses[index].Location.City
				state = restaurantResult.Businesses[index].Location.State
				country = restaurantResult.Businesses[index].Location.Country
			}
		}
	}

	// Creating destination object of which we will return
	destination := &models.Destination{
		Location:      [3]string{city, state, country},
		Restaurants:   RestaurantList,
		Entertainment: EntertainmentList,
		Shopping:      ShoppingList,
		Start:         params["start"],
		End:           params["end"],
	}

	// Returning destination object
	json.NewEncoder(w).Encode(destination)
}
