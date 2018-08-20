package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/web"
)

type PlanetHandler struct {
	repository repository.PlanetRepository
}

func NewPlanetHandler(repository repository.PlanetRepository) *PlanetHandler {
	return &PlanetHandler{repository}
}

func (p *PlanetHandler) Create(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)
	planet := &planet.Planet{}

	if err := json.Unmarshal(body, planet); err != nil {
		log.WithFields(log.Fields{"planet": planet, "err": err.Error()}).Error("invalid payload to Planet ")
		web.RespondError(w, err, http.StatusUnprocessableEntity)
		return
	}

	err := web.IsRequestValid(planet)
	if err != nil {
		web.RespondError(w, err, http.StatusBadRequest)
		return
	}

	if planet, err := p.repository.Store(planet); err != nil {
		log.WithFields(log.Fields{"planet": planet, "err": err.Error()}).Error("Error to save Planet.")
		web.RespondError(w, err, http.StatusInternalServerError)
	} else {
		log.WithFields(log.Fields{"planet": planet}).Info("planet save")
		web.Respond(w, planet, http.StatusCreated)

	}
}
