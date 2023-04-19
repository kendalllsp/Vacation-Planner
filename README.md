
# CEN3031: Travel Planner

We are building a Travel Planner that allows the user to type in where they want to go, and what month of the year they're planning on going. 
Then our program will output the recent weather trends over the last couple years, 
local food/entertainment spots with the best reviews, etc. 
If results haven’t been determined for that destination and time frame, 
we will fetch the info and input it in the database. 
If it has already been searched (destination and month) then we can search the database for the already fetched information. Using OpenWeatherAPI, YelpAI, and more(TBD), we will be able to create a seamless interface for our travel-ready users.


## Status
![Open Issues](https://img.shields.io/github/issues/leonleonardo/cen3031) ![Closed Issues](https://img.shields.io/github/issues-closed/leonleonardo/cen3031) ![Open Pull Requests](https://img.shields.io/github/issues-pr/leonleonardo/cen3031)

## Authors

#### Backend team

- [Benjamin Mendoza](https://www.github.com/benmendoza3)
- [Kendall Stansfield-Phillips](https://www.github.com/kendalllsp)

#### Frontend team
- [Richard Cusolito](https://www.github.com/rickcuso88)
- [Leonardo Leon](https://www.github.com/leonleonardo)

## Development

#### Dependencies

You need to have [Go](https://golang.org/),
[Node.js](https://nodejs.org/) and,
[Docker](https://www.docker.com/).

Verify the tools are installed by running the following commands:

```zsh
go version
npm --version
docker -v
```

Navigate to the `server` folder 

```zsh
cd server
```

Run docker compose to setup server and hot reload

```zsh
docker compose up
``` 

The back end will serve on http://localhost:8181.


#### Initializing the frontend

Navigate to the `webapp` folder

```sh
cd webapp
```
Install the dependencies  

```sh
npm install
```

Start the frontend server

```sh
npm start
```
The application will be available on http://localhost:4200.


