package enviroment

import (
	"fmt"
	"os"
)

type service struct {
	Port string `json:"port"`
}

type db struct {
	MysqlWrite string `json:"mysqlwrite"`
}

type config struct {
	Service service `json:"service"`
	Db      db      `json:"db"`
}

var Conf config

func Load() {
	var PORT_ENV string = os.Getenv("PORT_ENV")
	var MYSQL_WRITE string = os.Getenv("MYSQL_WRITE")

	if PORT_ENV == "" || MYSQL_WRITE == "" {
		panic(fmt.Errorf("Não foram encontradas as variáveis de ambiente para inicialização do sistema. Verifique a váriaveis 'PORT_ENV', 'MYSQL_WRITE'."))
	}

	Conf = config{service{PORT_ENV}, db{MYSQL_WRITE}}
}
