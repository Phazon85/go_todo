package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phazon85/go_todo/handler"
	"github.com/phazon85/go_todo/services"

	_ "github.com/lib/pq"
)

const (
	configFile = "dev.yaml"
)

//Todo holds json fields
type Todo struct {
	ID    string `json:"ID"`
	Title string `json:"Title"`
	Body  string `json:"Body"`
}

func main() {
	database := services.DBInit(configFile)
	handler := handler.NewTodoHandler(database)

	r := mux.NewRouter()
	r.HandleFunc("/", handler.HandleGetTodos).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// func (api *API) getTodo(w http.ResponseWriter, r *http.Request) {

// func (api *API) postTodo(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Incoming POST request on: %s", r.URL.Path)
// 	newTodo := &todo{}
// 	err := json.NewDecoder(r.Body).Decode(newTodo)
// 	if err != nil {
// 		log.Printf("Error decoding post Todo: %s", err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	sqlStatement := `
// 	INSERT INTO todo_list (title, body) VALUES ($1, $2);`

// 	_, err = api.DB.Exec(sqlStatement, newTodo.Title, newTodo.Body)
// 	if err != nil {
// 		log.Printf("Error performing INSERT statement: %s", err.Error())
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }

// func (api *API) delTodo(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Incoming DELETE request on: %s", r.URL.Path)
// 	deleteStatment := `
// 	DELETE FROM todo_list WHERE id = $1;`
// 	value := r.Header.Get("ID")
// 	_, err := api.DB.Exec(deleteStatment, value)
// 	if err != nil {
// 		log.Printf("Error deleting record: %s", err.Error())
// 	}
// 	w.WriteHeader(http.StatusOK)
// }

// func (api *API) putTodo(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("Incoming PUT request on: %s", r.URL.Path)
// 	updateStatement := `
// 	UPDATE todo_list SET title = $2, body = $3 WHERE id = $1;`
// 	updateTodo := &todo{}
// 	err := json.NewDecoder(r.Body).Decode(updateTodo)
// 	if err != nil {
// 		log.Printf("Error decoding put Todo: %s", err.Error())
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	_, err = api.DB.Exec(updateStatement, r.Header.Get("ID"), &updateTodo.Title, &updateTodo.Body)
// 	if err != nil {
// 		log.Printf("Error updating Todo: %s", err.Error())
// 	}

// 	w.WriteHeader(http.StatusOK)
// }
