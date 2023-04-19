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

// 1. Obtain private API key.
// 2. Refer to the request parameters for the location the user is attempting to search.
// 3. Initialize connect to the Yelp Fusion API.
// 4. Set relevant filters for Fashion business search.
// 5. Perform business search for Fashion businesses and check for any errors during Fashion business search.
// 6. Initialize variable and loop through results of Fashion Business search to save data needed for user.
// 7. Set relevant filters for Entertainment business search.
// 8. Perform business search for Entertainment businesses and check for any errors during Entertainment business search.
// 9. Initialize variable and loop through results of Entertainment Business search to save data needed for user.
// 10. Set relevant filters for Food business search.
// 11. Perform business search for Food businesses and check for any errors during Food business search.
// 12. Initialize variable and loop through results of Food Business search to save data needed for user.
// 13. Use result closest to location specified by user to determine returned Location values (city, state, country)
// 14. Initialize variable to store all of the search results, the location result derived above, and the start and end dates specified by the user.
// 15. Write filled response variable for front end.

func (h DBRouter) GetDestInfo(w http.ResponseWriter, r *http.Request) {

	// 1.
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file")
	}
	yelpAPIKey := os.Getenv("YELP_REST_API_KEY")

	// 2.
	params := mux.Vars(r)
	destinationLocation := params["location"]

	// 3.
	yelp := yfusion.NewYelpFusion(yelpAPIKey)

	// 4.
	businessSearch := &yfusion.BusinessSearchParams{}
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("fashion")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// 5.
	shoppingResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		fmt.Println("Fashion clothing business search could not be completed.")
	}

	// 6.
	var ShoppingList [10]models.Business
	for i, b := range shoppingResult.Businesses {
		ShoppingList[i].Name = b.Name
		ShoppingList[i].Price = b.Price
		ShoppingList[i].Rating = b.Rating
		ShoppingList[i].Location = b.Location.Address1
		ShoppingList[i].Type = b.Categories[0].Title

		// "Not listed" when the price is empty
		if len(ShoppingList[i].Price) < 1 {
			ShoppingList[i].Price = "Not listed"
		}
	}

	// 7.
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("arts")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// 8.
	entertainmentResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		log.Fatal(err)
	}

	// 9.
	var EntertainmentList [10]models.Business
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

	// 10.
	businessSearch.SetLocation(destinationLocation)
	businessSearch.SetTerm("food")
	businessSearch.SetLimit(10)
	businessSearch.SetRadius(15000)
	businessSearch.SetSortBy("rating")

	// 11.
	restaurantResult, err := yelp.SearchBusiness(businessSearch)
	if err != nil {
		log.Fatal(err)
	}

	// 12.
	var RestaurantList [10]models.Business
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

	// 13.
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

	// 14.
	destination := &models.Destination{
		Location:      [3]string{city, state, country},
		Restaurants:   RestaurantList,
		Entertainment: EntertainmentList,
		Shopping:      ShoppingList,
		Start:         params["start"],
		End:           params["end"],
	}

	// 15.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(destination)
}
