package web

import (
	"github.com/gorilla/mux"
)

func NewAPI() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	api := r.PathPrefix("/api/v1").Subrouter()

	return api
}
