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
	psqlInfo   = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
)

//PSQLService implements the Actions interface and carries the sql connection info
type PSQLService struct {
	DB *sql.DB
}

const (
	allTodos   = "SELECT id, title, body FROM todo_list;"
	todoByID   = "SELECT id, title, body FROM todo_list WHERE id=$1;"
	createTodo = "INSERT INTO todo_list (title, body) VALUES ($1, $2)"
	deleteTodo = "DELETE FROM todo_list WHERE id = $1;"
	updateTodo = "UPDATE todo_list SET title = $2, body = $3 WHERE id = $1;"
)

// DBInit takes a config struct and returns a postgres DB connection
func DBInit(file string) *PSQLService {
	cfg := config.NewConfig(file)
	s := cfg.Service
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

//AllTodos Gets all Todos from DB
func (s *PSQLService) AllTodos() ([]*Todo, error) {
	allTodo := []*Todo{}
	rows, err := s.DB.Query(allTodos)
	if err != nil {
		log.Printf("Error with allTodos query")
	}
	defer rows.Close()
	for rows.Next() {
		newTodo := &Todo{}
		err = rows.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body)
		if err != nil {
			log.Printf("Error scanning todos to allTodos: %s", err.Error())
		}
		allTodo = append(allTodo, newTodo)
	}
	return allTodo, err
}

//GetTodoByID gets single Todo by ID from DB
func (s *PSQLService) GetTodoByID(id string) (*Todo, error) {
	newTodo := &Todo{}
	if id == "" {
		return nil, ErrNoID
	}
	row := s.DB.QueryRow(todoByID, id)
	switch err := row.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body); err {
	case sql.ErrNoRows:
		return nil, err
	case nil:
		return newTodo, nil
	default:
		return nil, err
	}
}

//AddTodo adds a new todo into the DB
func (s *PSQLService) AddTodo(todo *Todo) error {
	_, err := s.DB.Exec(createTodo, todo.Title, todo.Body)
	return err
}

//UpdateTodo updates specified Todo by ID
func (s *PSQLService) UpdateTodo(todo *Todo) error {
	_, err := s.DB.Exec(updateTodo, todo.ID, todo.Title, todo.Body)
	return err
}

//DeleteTodo deletes todo by id
func (s *PSQLService) DeleteTodo(id string) error {
	_, err := s.DB.Exec(deleteTodo, id)
	return err
}
