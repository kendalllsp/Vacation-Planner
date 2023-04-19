
# CEN3031: Travel Planner

Welcome to Travel Planner a web application that built to make planning vacations easy. The main search bar allows the user to look up where they want to go, along with their prefered dates of travel. 
Then Travel Planner will output the local restaurants, shopping, and entertainment spots with the best reviews. 
Our selection of local spots is based on a filtered list pulled from the YelpAPI. Users can save destinations and see their different trips in one dashboard. 

## Status
![Open Issues](https://img.shields.io/github/issues/leonleonardo/cen3031) ![Closed Issues](https://img.shields.io/github/issues-closed/leonleonardo/cen3031) ![Open Pull Requests](https://img.shields.io/github/issues-pr/leonleonardo/cen3031)

## Authors

#### Backend team

- [Benjamin Mendoza](https://www.github.com/benmendoza3)
- [Kendall Stansfield-Phillips](https://www.github.com/kendalllsp)

#### Frontend team
- [Richard Cusolito](https://www.github.com/rickcuso88)
- [Leonardo Leon](https://www.github.com/leonleonardo)

## Screenshots



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
#### Initializing backend

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


