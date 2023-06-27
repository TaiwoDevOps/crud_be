package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // imported but not used and escaped by underscore
	"github.com/rs/zerolog/log"
	"taiwodevops.com/crud_be/helper"
)

// setup DB config const val

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbName   = "crudOps"
)

// function to establish connection to DB
func DatabaseConnection() *sql.DB {

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	//open the db connection
	db, err := sql.Open("postgres", sqlInfo)

	helper.PanicIfError(err)

	err = db.Ping()
	helper.PanicIfError(err)

	log.Info().Msg("Connected to Database")

	return db

}
