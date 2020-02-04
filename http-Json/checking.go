package main

import "log"

func checkPersonById(id1, id2 string) bool {
	return id1 == id2
}

func checkError(err error) {
	if err != nil {
		log.Print(err)
	}
}