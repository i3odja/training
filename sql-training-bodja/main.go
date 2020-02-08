package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	address := flag.String("port", ":8099", "port")
	flag.Parse()

	sp := mewSQLPeople()
	defer sp.dbase.Close()

	mainRoute := sp.Routing()

	fmt.Printf("[port%v] Server Listening...", *address)
	err := http.ListenAndServe(*address, mainRoute)
	if err != nil {
		log.Fatal(err.Error())
	}
}
