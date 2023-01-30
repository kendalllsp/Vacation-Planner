package main

type Business struct {
    Name        string  `json: "name"`
    Price       string  `json: "price"`
    Rating      float64 `json: "rating"`
    Location    string  `json: "location"`
    Type        string  `json: "type"`
}

type Destination struct {
    Location        [3]string        `json: "location"`
    Restaurants     [10]Business    `json: "restaurants"`
    Entertainment   [10]Business    `json: "entertainment"`
    Shopping        [10]Business    `json: "shopping"`
}