package config

import (
	"database/sql"
	"fmt"
	"github.com/rs/zerolog/log"
	"golang-crud/Helpers"
    _"github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbNamw   = "test"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s user=%s", host, user)
	db, err := sql.Open("postgres", sqlInfo)
	helpers.PanicIfError(err)

	err = db.Ping()
	helpers.PanicIfError(err)

	log.Info().Msg("Connected to database")

	return db

}
