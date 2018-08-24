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

With Golang:

```sh
  go install /cmd/swplanetsd/main.go
```

With Docker:

```sh
  docker-compose up -d --build
```

## Documentation and Examples
