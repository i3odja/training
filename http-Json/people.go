package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	jsonIterator "github.com/json-iterator/go"
)

type Person struct {
	ID      	string `json:"ID"`
	FirstName 	string `json:"FirstName"`
	LastName	string `json:"LastName"`
	Age     	int    `json:"Age"`
}

type People struct {
	people []Person
}

// [C]reate[R][U][D]
func (pr *People) AddPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var person Person

	checkError(jsonIterator.NewDecoder(r.Body).Decode(&person))
	person.ID = uuid.New().String()
	pr.people = append(pr.people, person)
	checkError(jsonIterator.NewEncoder(w).Encode(&person))
}

// [C][R]ead[U][D]
func (pr *People) GetPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	checkError(jsonIterator.NewEncoder(w).Encode(pr.people))
}

// [C]reate[R][U][D]
func (pr *People) GetPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	for _, person := range pr.people {
		if checkPersonById(person.ID, vars["ID"]) {
			checkError(jsonIterator.NewEncoder(w).Encode(person))
			return
		}
	}
}

// [C][R][U]pdate[D]
func (pr *People) EditPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	for i, person := range pr.people {
		if checkPersonById(person.ID, vars["ID"]) {
			var p Person

			checkError(jsonIterator.NewDecoder(r.Body).Decode(&p))
			p.ID = person.ID
			checkError(jsonIterator.NewEncoder(w).Encode(p))
			pr.people[i] = p
		}
	}
}

// [C][R][U][D]elete
func (pr *People) DeletePersonFromPeople(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)

	for i, person := range pr.people {
		if checkPersonById(person.ID, vars["ID"]) {
			pr.people = append(pr.people[:i], pr.people[i+1:]...)
			break
		}
	}
}