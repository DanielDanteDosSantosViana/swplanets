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

type config struct {
	Service service
	Db      db
}

var Conf config

func Load() {
	var PORT_ENV string = os.Getenv("PORT_ENV")
	var MONGO_HOST string = os.Getenv("MONGO_HOST")
	var DB_NAME string = os.Getenv("DB_NAME")

	if PORT_ENV == "" || MONGO_HOST == "" || DB_NAME == "" {
		panic(fmt.Errorf("it was not possible to find the environment variables 'PORT_ENV', 'MONGO_HOST', 'DB_NAME'."))
	}

	Conf = config{service{PORT_ENV}, db{MONGO_HOST, DB_NAME}}
}
