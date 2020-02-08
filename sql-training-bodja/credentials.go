package main

import "os"

type Credentials struct {
	userName string
	password string
	dbname   string
}

func (dbCredentials *Credentials) SetCredentials() string {
	dbCredentials.userName = os.Getenv("MYSQL_USER")
	dbCredentials.password = os.Getenv("MYSQL_PASSWORD")
	dbCredentials.dbname = os.Getenv("MYSQL_DB")

	creds := dbCredentials.userName + ":" + dbCredentials.password + "@/" + dbCredentials.dbname
	return creds
}
