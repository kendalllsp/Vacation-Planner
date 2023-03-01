package routes

import (
    "net/http"
    "github.com/gorilla/mux"
    "encoding/json"
    "github.com/jmatth11/yfusion"
    "fmt"
    // Removed these two purposes of running tests
    // "os"
    // "github.com/joho/godotenv"
    "vacation-planner/models"
)

func (h DBRouter) GetDestInfo(w http.ResponseWriter, r *http.Request) {

    w.Header().Set("Content-Type", "application/json")

    // Using godotenv to hide API keys

    // COMMENTING OUT FOR TESTING PURPOSES
    // err := godotenv.Load()
    // if err != nil {
    //     log.Fatal("Could not load .env file")
    // }
    // COMMENTED OUT FOR TESTING PURPOSES

    // Setting key var to YELP
    // COMMENTED OUT FOR TESTING PURPOSES
    // yelpAPIKey := os.Getenv("YELP_REST_API_KEY")
    yelpAPIKey := "w09RoocygtRReubGF_vLBYqIszZkyzBukVeOtPo74PMAeQ9QzRmLGYF2WcE5KkhLuHVUWGAyFHSMeqQUfmmVgrHKe1kFDeqPbr53cpLcBtbUjKqYVlPTB770HJXUY3Yx"
    // COMMENT ^ OUT WHEN TRYING TO RUN PROGRAM OUTSIDE OF TESTS

    // Taking the location value typed by the user, to use for computations
    params := mux.Vars(r)
    destinationLocation := params["location"]

    // Starting new Yelp client
    yelp := yfusion.NewYelpFusion(yelpAPIKey)

    // Setting all relevant filters for specific search
    businessSearch := &yfusion.BusinessSearchParams{}
    businessSearch.SetLocation(destinationLocation)
    businessSearch.SetTerm("fashion")
    businessSearch.SetLimit(10)
    businessSearch.SetRadius(15000)
    businessSearch.SetSortBy("rating")
    shoppingResult, err := yelp.SearchBusiness(businessSearch)
    if err != nil {
        fmt.Println("Fashion clothing business search could not be completed.")
    }

    // Creating new slice of size to for the top 10 rated Shopping businesses
    var ShoppingList [10]models.Business

    // Looping through results of the search to populate slice of shopping businesses
    for i, b := range shoppingResult.Businesses {
        ShoppingList[i].Name = b.Name
        ShoppingList[i].Price = b.Price
        ShoppingList[i].Rating = b.Rating
        ShoppingList[i].Location = b.Location.Address1
        ShoppingList[i].Type = b.Categories[0].Title

        // Filling price string with not listed rather than empty string
        if (len(ShoppingList[i].Price) < 1) {
            ShoppingList[i].Price = "Not listed"
        }
    }

    // Entertainment search
    businessSearch.SetLocation(destinationLocation)
    businessSearch.SetTerm("arts")
    businessSearch.SetLimit(10)
    businessSearch.SetRadius(15000)
    businessSearch.SetSortBy("rating")
    entertainmentResult, err := yelp.SearchBusiness(businessSearch)
    if err != nil {
        fmt.Println("X")
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

        if (len(EntertainmentList[i].Price) < 1) {
            EntertainmentList[i].Price = "Not listed"
        }
    }

    // Food search
    businessSearch.SetLocation(destinationLocation)
    businessSearch.SetTerm("food")
    businessSearch.SetLimit(10)
    businessSearch.SetRadius(15000)
    businessSearch.SetSortBy("rating")
    restaurantResult, err := yelp.SearchBusiness(businessSearch)
    if err != nil {
        fmt.Println("X")
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

        if (len(RestaurantList[i].Price) < 1) {
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
        Location: [3]string{city, state, country},
        Restaurants: RestaurantList,
        Entertainment: EntertainmentList,
        Shopping: ShoppingList,
    }

    // // Decoding request
    // _ = json.NewDecoder(r.Body).Decode(&destination)

    // Returning destination object
    json.NewEncoder(w).Encode(destination)

}