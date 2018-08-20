package main

import (
	"net/http"
	"os"

	"github.com/DanielDanteDosSantosViana/swplanets/cmd/swplanetsd/planet/web/route"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/enviroment"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/web"

	"github.com/rs/cors"
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func main() {
	enviroment.Load()

	session, err := db.NewSession()
	if err != nil {
		log.Panicf(err.Error())
	}

	api := web.NewAPI()
	route.AddAPI(session, api)

	log.Info(" SWPlanets running on port ", enviroment.Conf.Service.Port)
	handler := cors.Default().Handler(api)

	err = http.ListenAndServe(":"+enviroment.Conf.Service.Port, handler)
	if err != nil {
		log.Fatal("Error init Server : ", err)
	}

}
