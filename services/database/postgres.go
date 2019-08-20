package database

import (
	"database/sql"
	"fmt"

	"github.com/Phazon85/go_todo/services/config"
)

// setting up connection to postgres DB
// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 	host, port, user, password, dbname)
// db, err := sql.Open("postgres", psqlInfo)
// if err != nil {
// 	log.Printf("Error opening SQL db: %s", err.Error())
// }
// err = db.Ping()
// if err != nil {
// 	log.Printf("Error pingng SQL db: %s", err.Error())
// }
// api := &API{
// 	DB: db,
// }

// DBInit takes a config struct and returns a postgres DB connection
func DBInit(driver string, config *config.Config) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s ssqlmode=disable", config.host, config.port, config.user, config.password, config.dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Prinf("Error opening SQL db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pingng SQL db: %s", err.Error())
	}
	return db
}
