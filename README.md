<p align="center">
  <img alt="SwPlanets" src="https://pbs.twimg.com/media/Ca4yzw7WwAMlTqw.png" height="140" />
</p>

# SWPlanets
API to create planets with the quantities of appearances in Star Wars movies


## How to Install

Requirements:

  * Docker
    [docker](https://www.docker.com/).
  * Dep: dependencies manager , see the Software section in the
    [dep_golang](https://github.com/golang/dep).

Install dependencies:

```sh
  dep ensure
```
## Run

Enviroment:

SWplanets use mongodb to save date and get data about star wars movies in external API (https://swapi.co/api/).
Database connections and external api's url are defined with environments variables:

* PORT_ENV: api port
* DB_NAME: name of database
* MONGO_HOST: url connection database
* URL_API: database  connection  url 

Test:

```sh
  make test
```

Golang:

```sh
  make run
```

Docker:

```sh
  docker-compose up -d --build
```

## Documentation and Examples

Endpoints:

  * CREATE PLANET - /api/v1/planets (POST) 
  * GET PLANET BY ID - /api/v1/planets/{id} (GET) 
  * GET PLANET BY NAME - /api/v1/planets?name={name} (GET)
  * LIST PLANET - /api/v1/planets (GET)
  * REMOVE PLANET - /api/v1/planets/{id} (DELETE)
  
  
Payload:

```json
 {

    "name":"Alderaan",
    "climate":"temperate",
    "terrain":"gas giant"
  }
```
 Example with curl:
 
 
```sh
curl -X POST \
  http://localhost:8081/api/v1/planets \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -d '{
	"name":"Alderaan",
	"climate":"temperate",
	"terrain":"gas giant"
}'

```

