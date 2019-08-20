package postgres

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/phazon85/go_todo/services/config"
)

const (
	driverName = "postgres"
	psqlInfo   = "host=%s port=%d user=%s password=%s dbname=%s ssqlmode=disable"
)

// DBInit takes a config struct and returns a postgres DB connection
func DBInit(config *config.Config) *sql.DB {
	s := config.Service
	psql := fmt.Sprintf(psqlInfo, s.Host, s.Port, s.User, s.Password, s.Name)
	db, err := sql.Open(driverName, psql)
	if err != nil {
		log.Printf("Error opening SQL db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pingng SQL db: %s", err.Error())
	}
	return db
}
