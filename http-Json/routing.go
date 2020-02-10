package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (sp *SQLPeople) Routing() *mux.Router {
	mainRoute := mux.NewRouter()
	mainRoute.Use(Middleware)
	apiRoute := mainRoute.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/list/page/{PAGE}", sp.GetAllPeople).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list/{ID}", sp.GetOnePerson).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list", sp.AddNewPerson).Methods(http.MethodPost)
	apiRoute.HandleFunc("/list/{ID}", sp.UpdateOnePerson).Methods(http.MethodPut)
	apiRoute.HandleFunc("/list/{ID}", sp.DeleteOnePerson).Methods(http.MethodDelete)

	return mainRoute
}
