package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	address := flag.String("a", ":8099", "url port")
	flag.Parse()

	var pr People

	mainRoute := mux.NewRouter()
	apiRoute := mainRoute.PathPrefix("/api/v1").Subrouter()
	// [C]reat-[R]ead-[U]pdate-[D]elete
	apiRoute.HandleFunc("/list", pr.GetPeople).Methods("GET")
	apiRoute.HandleFunc("/list/{ID}", pr.GetPerson).Methods("GET")
	apiRoute.HandleFunc("/list", pr.AddPerson).Methods("POST")
	apiRoute.HandleFunc("/list/{ID}", pr.EditPerson).Methods("PUT")
	apiRoute.HandleFunc("/list/{ID}", pr.DeletePersonFromPeople).Methods("DELETE")

	fmt.Println("Server Listening...")
	err := http.ListenAndServe(*address, mainRoute)
	if err != nil {
		log.Fatal(err.Error())
	}
}
