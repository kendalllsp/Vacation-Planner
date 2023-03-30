# **Work Completed for Sprint 3:**

- Created tests for frontend
- Created 'logged in' status variable to detect when a user is logged in even when changing pages
- Added ability for user to save and delete locations
- Added ability for user to see list of saved destions
- Created new unit testing for the added backend functionality
- Created new UI for search results page

# **Unit Tests:**

#### Frontend:

#### Backend:
- create user function test
- login user function test
- get destination function test
- database connection test

# **Updated Backend Documentation:**

## Overview:
This API is designed to store and access user login information in a database and contact a Yelp API to return information about travel destinations, sorting through the Yelp API output to return only relevant information. In addition, users can save locations which they are interested in, and may view and edit this list by adding to it or deleting from it. Generally, this API provides users the ability to search for new travel destinations and save their favorites for future reference.

## Dependencies: 
The only additional dependency this API requires for full functionality is a .env file containing a private Yelp API key. A contributor must be contacted regarding this for security reasons.

## Yelp API integration:
This API integrates with the Yelp API to provide information to users about shopping, restaurants, and entertainment near a searched destination. Currently, the top 10 most highly rated restaurants, stores, and entertainment locations are returned, displaying their rating along with the average price, address, and the type of establishment (i.e. "French Restaurant").

## Endpoints:
### **POST** /createUser - Creates a new user account

Parameters:
- Request body is a JSON file in the following format:
{
Email: "email@example.com",
Password : "password"
}

Response: 
- Response is a JSON file in the following format:
{ Success: true, Message: "User succesfully created account" }



### **POST** /loginUser - Authenticates a user login

Parameters:
- Request body is a JSON file in the following format:
{
Email: "email@example.com",
Password : "password"
}

Response: 
- Response is a JSON file in the following format:
{ LoggedIn: true, Email: "email@example.com", Message: "User successfully logged in." }

- LoggedIn value will register as false if the email is unrecognized or the email and password do not match.

### **GET** /newDestination/{location} - Returns Yelp API information about a specific location

Parameters:
- Function takes in the parameter 'location' from the URL

Response:
- Response is a JSON file containing a destination object, which has the following format:

{

        Location: [3]string{city, state, country},
        
        Restaurants: RestaurantList,
        
        Entertainment: EntertainmentList,
        
        Shopping: ShoppingList,
}

- Each 'list' is of length 10, and contains the names of 10 locations along with their ratings, address, type, and price.


## Endpoints under development:

### **GET** /updateDestination - Returns all saved destinations

Parameters: 
- Request body is a JSON file in the following format:
{
Email: "email@example.com",
}

Response:
- Function response is a byte slice that lists the users saved locations if it is populated, "User destination list is empty." if it is not, or states "No user with the email address associated." if both other conditions are false.

### **PUT** /updateDestination - Saves a new location in the database

Parameters: 
- Request body is a JSON file in the following format:
{
Email: "email@example.com",
Location : "exampleLocation"
}

Response:
- Function response is a byte slice, either stating "New location successfully saved." or "Location already saved."


### **DELETE** /updateDestination - Deletes a user's saved location

Parameters: 
- Request body is a JSON file in the following format:
{
Email: "email@example.com",
Location : "exampleLocation"
}

Response:
- Function response is a byte slice, either stating "Location successfully deleted." or "Account does not have location saved in order to be deleted."



# **Demo Video:**
This is a link to a video detailing the current functionality of our integrated application:

