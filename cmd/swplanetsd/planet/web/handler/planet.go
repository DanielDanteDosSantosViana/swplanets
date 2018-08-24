package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/planet/repository"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/client"
	"github.com/DanielDanteDosSantosViana/swplanets/internal/platform/web"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

type PlanetHandler struct {
	repository repository.PlanetRepository
	swApi      client.Client
}

func NewPlanetHandler(planetRepository repository.PlanetRepository, client client.Client) *PlanetHandler {
	return &PlanetHandler{planetRepository, client}
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

	appearances, err := p.swApi.GetNumberOfAppearancesByPlanetName(planet.Name)
	if err != nil {
		switch err.(type) {
		case *client.BadMsgError:
			log.WithFields(log.Fields{"planet": planet, "err": err.Error()}).Error("Error when send data to external API.")
			web.RespondError(w, err, err.(*client.BadMsgError).StatusCode())
			return
		default:
			log.WithFields(log.Fields{"planet": planet}).Error("Internal error when get data about planet in external api")
			web.RespondError(w, err, http.StatusBadGateway)
			return
		}
	}

	planet.Appearances = appearances

	if planet, err := p.repository.Store(planet); err != nil {
		switch err.(type) {
		case *repository.DuplicationKeyError:
			log.WithFields(log.Fields{"planet": planet}).Error(err.Error())
			web.RespondError(w, err, http.StatusConflict)
		default:
			log.WithFields(log.Fields{"planet": planet}).Error("Internal error when insert planet")
			web.RespondError(w, err, http.StatusInternalServerError)
		}
	} else {
		log.WithFields(log.Fields{"planet": planet}).Info("planet save")
		web.Respond(w, planet, http.StatusCreated)

	}
}

func (p *PlanetHandler) Remove(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParams := vars["id"]

	if err := p.repository.Remove(idParams); err != nil {

		switch err.(type) {
		case *repository.InvalidIdError:
			log.WithFields(log.Fields{"id": idParams}).Error(err.Error())
			web.RespondError(w, err, http.StatusBadRequest)
		default:
			log.WithFields(log.Fields{"id": idParams}).Error(err.Error())
			web.RespondError(w, err, http.StatusInternalServerError)
		}

	} else {
		web.Respond(w, nil, http.StatusNoContent)
	}

}

func (p *PlanetHandler) List(w http.ResponseWriter, r *http.Request) {

	if planets, err := p.repository.List(); err != nil {
		log.Error(err.Error())
		web.RespondError(w, err, http.StatusInternalServerError)
	} else {
		web.Respond(w, planets, http.StatusOK)
	}
}

func (p *PlanetHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParams := vars["id"]

	if planet, err := p.repository.GetById(idParams); err != nil {

		switch err.(type) {
		case *repository.NotFoundError:
			log.WithFields(log.Fields{"planet id": idParams}).Error(err.Error())
			web.RespondError(w, err, http.StatusNotFound)
		default:
			log.WithFields(log.Fields{"id": idParams}).Error(err.Error())
			web.RespondError(w, err, http.StatusInternalServerError)
		}

	} else {
		web.Respond(w, planet, http.StatusOK)
	}

}
