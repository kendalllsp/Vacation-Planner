# **Work Completed for Sprint 4:**

- Created new tests for frontend
- Added ability for user to see businesses in saved trips in UI
- Added ability to save and display trip dates in UI
- Created new unit testing for the added backend functionality
- Added ability to save individual businesses to a database (not available in UI)

# **Unit Tests:**

#### Frontend:
- home component test
- login component test
- signup component test
- navbar component test
- trip component and subcomponent tests
- navbar routing test 
- login routing test
- home page routing test
- trips page routing test
- routing tests to originate from each page

#### Backend:
- create user function test
- login user function test
- get destination function test
- database connection test
- update destination test
- update saved businesses test

# **Updated Backend Documentation:**

## Overview:
This API is designed to store and access user login information in a database and contact a Yelp API to return information about travel destinations, sorting through the Yelp API output to return only relevant information. In addition, users can save locations which they are interested in, and they may view and edit this list by adding to it or deleting from it. Generally, this API provides users the ability to search for new travel destinations and save their favorites for future reference.

## Dependencies: 
The only additional dependency this API requires for full functionality is a .env file containing a private Yelp API key and private database information. A contributor must be contacted regarding this for security reasons.

## Yelp API integration:
This API integrates with the Yelp API to provide information to users about shopping, restaurants, and entertainment businesses near a searched destination. Currently, the top 10 most highly rated restaurants, stores, and entertainment locations are returned, displaying their rating along with the average price, address, and the type of establishment (i.e. "French Restaurant").

## Endpoints:
### **POST** /createUser - Creates a new user account

Parameters:
- Request body is a JSON file in the following format:
{

        Email: "email@example.com",
        
        Password: "password"
}

Response: 
- Response is a JSON file in the following format:
{ Success: true, Message: "User succesfully created account" }


### **POST** /loginUser - Authenticates a user login

Parameters:
- Request body is a JSON file in the following format:
{

        Email: "email@example.com",
        
        Password: "password"
}

Response: 
- Response is a JSON file in the following format:
{ LoggedIn: true, Email: "email@example.com", Message: "User successfully logged in." }

- LoggedIn value will register as false if the email is unrecognized or the email and password do not match.

### **GET** /newDestination/{location}/ - Returns Yelp API information about a specific location

Parameters:
- Function takes in the parameter 'location' as well as 'start' and 'end' from the URL like so:
/newDestination?location=Austin,TX&start=04/01/2023&end=04/04/2023
where the parameters would be location = Austin, TX, start = 04/01/2023, and end would be 04/04/2023

Response:
- Response is a JSON file containing a destination object, which has the following format:

{

        Location: [3]string{city, state, country},
        
        Restaurants: RestaurantList,
        
        Entertainment: EntertainmentList,
        
        Shopping: ShoppingList,
        
        Start:         params["start"],
        
        End:           params["end"],
}

- Each 'list' is of length 10, and contains the names of 10 locations along with their ratings, address, type, and price.


### **GET** /updateDestination - Returns all saved destinations

Parameters: 
- Function takes in user's email address, as a query parameter in the URL like so:
 /updateDestination?Email=email@example.com

Response:
- Function returns a JSON file with all locations stored in an array, where each location is in the format:
{

        Email: "email@example.com",
        
        Location: "exampleLocation"
        
        Start: "2023-01-01"
        
        End: "2023-01-30"
}

### **POST** /updateDestination - Saves a new location in the database

Parameters: 
- Request body is a JSON file in the following format:
{

        Email: "email@example.com",
        
        Location: "exampleLocation"
        
        Start: "2023-01-01"
        
        End: "2023-01-30"
}

Response:
- Response is a JSON file in the following format:
{ Saved: true, Message: "New location successfully saved." }


### **DELETE** /updateDestination - Deletes a user's saved location

Parameters: 
- Request body is a JSON file in the following format:
{

        Email: "email@example.com",
        
        Location: "exampleLocation"
        
        Start: "2023-01-01"
        
        End: "2023-01-30"
}

Response:
- Response is a JSON file in the following format:
{ Saved: false, Message: "Location successfuly deleted." }


### **GET** /updateBusinessList - Returns all saved destinations

Parameters: 
- Function takes in user's email address and the location we want businesses from as query parameters in the URL like so:
 /updateBusinessList?email=email@example.com&location=Austin,TX

Response:
- Function returns a JSON file with all businesses stored in an array, where each location is in the format:
{

        "Name": "Jim's BBQ",
        
        "Price": "$",
        
        "Rating": 5.5,
        
        "B_Location": "123 Example Dr.",
        
        "Type": "BBQ"

}

### **POST** /updateBusinessList - Saves a new location in the database

Parameters: 
- Request body is a JSON file in the following format:
{

        "Email": "123@gmail.com",
        
        "Location": "Austin, TX",
        
        "Name": "Jim's BBQ",
        
        "Price": "$",
        
        "Rating": 5.5,
        
        "B_Location": "123 Example Dr.",
        
        "Type": "BBQ"
}

Response:
- Response is a JSON file in the following format:
{ Saved: true, Message: "New location successfully saved." }


### **DELETE** /updateBusinessList - Deletes a user's saved location

Parameters: 
- Request body is a JSON file in the following format:
{

        "Email": "123@gmail.com",
        
        "Location": "Austin, TX",
        
        "Name": "Jim's BBQ",
        
        "Price": "$",
        
        "Rating": 5.5,
        
        "B_Location": "123 Example Dr.",
        
        "Type": "BBQ"
}

Response:
- Response is a JSON file in the following format:
{ Saved: false, Message: "Location successfuly deleted." }

# **Demo Video:**
This is a link to a video detailing the final functionality of our integrated application:
[https://www.canva.com/design/DAFgmJHBM2E/watch](https://www.canva.com/design/DAFgmJHBM2E/daAslWTDXyd5vmHesiuI9g/watch?utm_content=D[…]t_publish&utm_medium=link&utm_source=celebratory_first_publish)
