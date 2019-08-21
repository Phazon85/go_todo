package services

import (
	"database/sql"
	"fmt"
	"log"

	//used for testing in postgres_test.go
	_ "github.com/lib/pq"
	"github.com/phazon85/go_todo/services/config"
)

const (
	driverName = "postgres"
	psqlInfo   = "host=%s port=%d user=%s password=%s dbname=%s ssqlmode=disable"
)

//PSQLService contains the sql DB connection info
type PSQLService struct {
	DB *sql.DB
}

// const (
// 	allTodos   = "SELECT id, title, body FROM todo_list;"
// 	todoByID   = "SELECT id, title, body FROM todo_list WHERE id=$1;"
// 	createTodo = "INSERT INTO todo_list (title, body) VALUES ($1, $2)"
// 	deleteTodo = "DELETE FROM todo_list WHERE id = $1;"
// 	updateTodo = "UPDATE todo_list SET title = $2, body = $3 WHERE id = $1;"
// )

// DBInit takes a config struct and returns a postgres DB connection
func DBInit(config *config.Config) *PSQLService {
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
	return &PSQLService{
		DB: db,
	}
}

func (s *PSQLService) AllTodos() ([]*Todo, error) {
	allTodo := []*Todo{}
	rows, err := s.DB.Query(allTodos)
	if err != nil {
		log.Printf("Error with allTodos query: %s", err.Error)
	}
	defer rows.Close()
	for rows.Next() {
		newTodo := &todo{}
		err = rows.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body)
		if err != nil {
			log.Printf("Error scanning todos to allTodos: %s", err.Error())
		}
		allTodo = append(allTodo, newTodo)
	}
	return allTodo, err
}

// 	// Get all todos
// 	allTodo := []*todo{}
// 	multiStatement := `SELECT id, title, body FROM todo_list;
// 		`
// 	rows, err := api.DB.Query(multiStatement)
// 	if err != nil {
// 		log.Printf("Error with GET multirow sql query: %s", err.Error())
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		newTodo := &todo{}
// 		err = rows.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body)
// 		if err != nil {
// 			log.Printf("Error scanning multi SQL rows into newTodo: %s", err.Error())
// 		}
// 		allTodo = append(allTodo, newTodo)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Printf("Error during iterating through sql reponse: %s", err.Error())
// 	}
// 	json.NewEncoder(w).Encode(allTodo)
// }
