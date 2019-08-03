package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Voltage13-2"
	dbname   = "todo"
)

// API holds postgres DB connection info
type API struct {
	DB *sql.DB
}

type todo struct {
	ID    string `json:"ID"`
	Body  string `json:"body"`
	Title string `json:"title"`
}

func (api *API) getTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming GET request on: %s", r.URL.Path)

	// // Single select statement
	// newTodo := &todo{}
	// value := r.Header.Get("ID")
	// singleStatment := `
	// SELECT id, title, body FROM todo_list WHERE id=$1;
	// `
	// row := api.DB.QueryRow(singleStatment, value)
	// switch err := row.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body); err {
	// case sql.ErrNoRows:
	// 	log.Printf("Error: No rows returned")
	// case nil:
	// 	fmt.Println(newTodo.ID, newTodo.Title, newTodo.Body)
	// default:
	// 	log.Printf("Error with GET sql query: %s", err.Error())
	// }

	// Get all todos
	allTodo := []*todo{}
	multiStatement := `SELECT id, title, body FROM todo_list;
		`
	rows, err := api.DB.Query(multiStatement)
	if err != nil {
		log.Printf("Error with GET multirow sql query: %s", err.Error())
	}
	defer rows.Close()
	for rows.Next() {
		newTodo := &todo{}
		err = rows.Scan(&newTodo.ID, &newTodo.Title, &newTodo.Body)
		if err != nil {
			log.Printf("Error scanning multi SQL rows into newTodo: %s", err.Error())
		}
		allTodo = append(allTodo, newTodo)
	}
	err = rows.Err()
	if err != nil {
		log.Printf("Error during iterating through sql reponse: %s", err.Error())
	}
	json.NewEncoder(w).Encode(allTodo)
}

func (api *API) postTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming POST request on: %s", r.URL.Path)
	newTodo := &todo{}
	err := json.NewDecoder(r.Body).Decode(newTodo)
	if err != nil {
		log.Printf("Error decoding post Todo: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	sqlStatement := `
	INSERT INTO todo_list (title, body) 
	VALUES ($1, $2);`

	_, err = api.DB.Exec(sqlStatement, newTodo.Title, newTodo.Body)
	if err != nil {
		log.Printf("Error performing INSERT statement: %s", err.Error())
	}
	w.WriteHeader(http.StatusCreated)
}

func (api *API) delTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming DELETE request on: %s", r.URL.Path)
	deleteStatment := `
	DELETE FROM todo_list
	WHERE id = $1;`
	value := r.Header.Get("ID")
	_, err := api.DB.Exec(deleteStatment, value)
	if err != nil {
		log.Printf("Error deleting record: %s", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func (api *API) putTodo(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming PUT request on: %s", r.URL.Path)
	updateStatement := `
	UPDATE todo_list
	SET title = $2, body = $3
	WHERE id = $1;`
	updateTodo := &todo{}
	err := json.NewDecoder(r.Body).Decode(updateTodo)
	if err != nil {
		log.Printf("Error decoding put Todo: %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
	}
	// for i, v := range todos {
	// 	if r.Header.Get("ID") == strconv.Itoa(v.ID) {
	// 		todos[i].Message = newTodo.Message
	// 	}
	// }
	_, err = api.DB.Exec(updateStatement, r.Header.Get("ID"), &updateTodo.Title, &updateTodo.Body)
	if err != nil {
		log.Printf("Error updating Todo: %s", err.Error())
	}

	w.WriteHeader(http.StatusOK)
}

func (api *API) rootHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		api.getTodo(w, r)
	case "POST":
		api.postTodo(w, r)
	case "DELETE":
		api.delTodo(w, r)
	case "PUT":
		api.putTodo(w, r)
	default:
		fmt.Fprintf(w, "Unknown method")
	}
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Printf("Error opening SQL db: %s", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Printf("Error pingng SQL db: %s", err.Error())
	}
	api := &API{
		DB: db,
	}
	http.HandleFunc("/todo", api.rootHandler)
	http.ListenAndServe(":8080", nil)
}
