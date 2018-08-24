package route

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/client"

	"github.com/DanielDanteDosSantosViana/swplanets/cmd/swplanetsd/planet/web/handler"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db"
	"github.com/gorilla/mux"
)

func AddAPI(sessiondb db.Session, api *mux.Router) *mux.Router {
	planetRepository := repository.NewMongoRepository(sessiondb)
	swApi := client.NewSWApi()
	planetHandler := handler.NewPlanetHandler(planetRepository, swApi)
	api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
	api.HandleFunc("/planets/{id}", planetHandler.Remove).Methods(http.MethodDelete)
	api.HandleFunc("/planets", planetHandler.List).Methods(http.MethodGet)
	api.HandleFunc("/planets/{id}", planetHandler.GetById).Methods(http.MethodGet)

	return api
}
