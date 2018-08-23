package enviroment

import (
	"fmt"
	"os"
)

type service struct {
	Port string
}

type db struct {
	Mongo string
	Name  string
}

type externalAPI struct {
	Url string
}

type config struct {
	Service     service
	Db          db
	ExternalAPI externalAPI
}

var Conf config

func Load() {
	var PORT_ENV string = os.Getenv("PORT_ENV")
	var MONGO_HOST string = os.Getenv("MONGO_HOST")
	var DB_NAME string = os.Getenv("DB_NAME")
	var URL_API string = os.Getenv("URL_API")

	if PORT_ENV == "" || MONGO_HOST == "" || DB_NAME == "" || URL_API == "" {
		panic(fmt.Errorf("it was not possible to find the environment variables 'PORT_ENV', 'MONGO_HOST', 'DB_NAME', 'URL_API'."))
	}

	Conf = config{service{PORT_ENV}, db{MONGO_HOST, DB_NAME}, externalAPI{Url: URL_API}}
}
