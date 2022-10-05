package component

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "database"
	port     = 5432 // Default port
	user     = "postgres"
	password = "docker_database"
	dbname   = "mytwitch"
)

var DB *sql.DB

func DBConnect() error {
	var err error
	DB, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname))
	if err != nil {
		return err
	}
	if err = DB.Ping(); err != nil {
		return err
	}
	return nil
}
