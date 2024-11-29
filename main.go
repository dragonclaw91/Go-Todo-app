package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/lib/pq"
)

// const creditScoreMin = 500
// const creditScoreMax = 900

type Task struct {
	Task        string `json:"task"`
	Id          string `json:"id"`
	IsCompleted string `json:"iscompleted"`
}

var db *sql.DB

const (
	host     = "localhost"
	port     = 5400
	user     = "postgres"
	password = "postgres"
	dbname   = "Todo"
)

func init() {
	var err error
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err = sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	// defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
}

func taskGet() ([]Task, error) {

	rows, err := db.Query(`SELECT * FROM "List" ORDER BY "isComplete"`)
	// err := db.Query()
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var tasks []Task
	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var task Task
		if err := rows.Scan(&task.Task, &task.Id, &task.IsCompleted); err != nil {
			return tasks, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return tasks, err
	}

	return tasks, nil

}

func taskPostPut(w http.ResponseWriter, r *http.Request, sql string, arg string) {

	// Handle the POST request...

	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	// Parse the JSON body
	var task Task
	if err := json.Unmarshal(body, &task); err != nil {
		http.Error(w, "Invalid JSON format", http.StatusBadRequest)
		return
	}
	sqlStatement := sql
	id := 0
	if arg == "POST" {
		err = db.QueryRow(sqlStatement, task.Task).Scan(&id)
	}
	if arg == "PUT" {
		err = db.QueryRow(sqlStatement, task.Id).Scan(&id)
	}
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", arg)

}

func handler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		taskGet()
	case http.MethodPost:
		taskPostPut(w, r, `
		INSERT INTO "List" ("task")
		VALUES ($1) RETURNING id`, "POST")
		fmt.Fprintln(w, "You made a POST request!")
	case http.MethodPut:
		taskPostPut(w, r, `UPDATE "List" SET "isComplete"=NOT "isComplete","completed_at"=NOW() WHERE "id"=$1 RETURNING id`, "PUT")
		fmt.Fprintln(w, "You made a PUT request!")
	case http.MethodDelete:
		fmt.Fprintln(w, "You made a DELETE request!")
		taskPostPut(w, r, ` DELETE FROM "List" WHERE "id"=$1 RETURNING id`, "PUT")
	default:
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}
}

func main() {
	// fmt.Println("New record ID is:", id)
	http.HandleFunc("/", handler)
	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
	fmt.Println("Successfully connected!")
}
