package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const offset = 10

type Person struct {
	ID        string `json:"ID"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Age       int    `json:"Age"`
}

type SQLPeople struct {
	dbase *sql.DB
}

func newSQLPeople() *SQLPeople {
	var cr Credentials

	credentials := cr.SetCredentials()
	dataBase, err := sql.Open("mysql", credentials)
	if err != nil {
		log.Println(err)
	}

	return &SQLPeople{
		dbase: dataBase,
	}
}

func (s *SQLPeople) AddNewPerson(w http.ResponseWriter, r *http.Request) {
	var person Person

	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	person.ID = uuid.New().String()

	_, err := s.dbase.Exec("INSERT INTO people(ID,Firstname,Lastname,Age) VALUES(?,?,?,?)", person.ID, person.Firstname, person.Lastname, person.Age)
	if err != nil {
		log.Println(err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *SQLPeople) GetAllPeople(w http.ResponseWriter, r *http.Request) {
	var people []Person

	vars := mux.Vars(r)
	page, err := strconv.Atoi(vars["PAGE"])

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	page--

	rows, err := s.dbase.Query("SELECT ID,Firstname,Lastname,Age FROM people WHERE Disabled=0 LIMIT ?,?", page*offset, offset)
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Person

		if err = rows.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Age); err != nil {
			log.Println(err)
			continue
		}

		people = append(people, p)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return
	}

	if len(people) == 0 {
		log.Println("Error: No Rows")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err = json.NewEncoder(w).Encode(people); err != nil {
		log.Println(err)
	}
}

func (s *SQLPeople) GetOnePerson(w http.ResponseWriter, r *http.Request) {
	var p Person

	vars := mux.Vars(r)

	row := s.dbase.QueryRow("SELECT ID,Firstname,Lastname,Age FROM people WHERE ID=?", vars["ID"])

	if err := row.Scan(&p.ID, &p.Firstname, &p.Lastname, &p.Age); err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(p); err != nil {
		log.Println(err.Error())
	}
}

func (s *SQLPeople) UpdateOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var p Person

	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p.ID = vars["ID"]

	if _, err := s.dbase.Exec("UPDATE people SET Firstname=?,Lastname=?,Age=? WHERE ID=?", p.Firstname, p.Lastname, p.Age, p.ID); err != nil {
		log.Println(err.Error())
		return
	}
}

func (s *SQLPeople) DeleteOnePerson(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, err := s.dbase.Exec("UPDATE people SET Disabled=1 WHERE ID=?", vars["ID"]); err != nil {
		log.Println(err.Error())
		return
	}
}
