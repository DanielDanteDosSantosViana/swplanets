package route

import (
	"net/http"

	"github.com/DanielDanteDosSantosViana/swplanets/cmd/swplanetsd/planet/web/handler"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/db"
	"github.com/gorilla/mux"
)

func AddAPI(sessiondb db.Session, api *mux.Router) *mux.Router {
	planetRepository := repository.NewMongoRepository(sessiondb)
	planetHandler := handler.NewPlanetHandler(planetRepository)
	api.HandleFunc("/planets", planetHandler.Create).Methods(http.MethodPost)
	return api
}
