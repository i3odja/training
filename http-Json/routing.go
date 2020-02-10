package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (s *SQLPeople) Routing() *mux.Router {
	mainRoute := mux.NewRouter()
	mainRoute.Use(Middleware)
	apiRoute := mainRoute.PathPrefix("/api/v1").Subrouter()
	apiRoute.HandleFunc("/list/page/{PAGE}", s.GetAllPeople).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list/{ID}", s.GetOnePerson).Methods(http.MethodGet)
	apiRoute.HandleFunc("/list", s.AddNewPerson).Methods(http.MethodPost)
	apiRoute.HandleFunc("/list/{ID}", s.UpdateOnePerson).Methods(http.MethodPut)
	apiRoute.HandleFunc("/list/{ID}", s.DeleteOnePerson).Methods(http.MethodDelete)

	return mainRoute
}
