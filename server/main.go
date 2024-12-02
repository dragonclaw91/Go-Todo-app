package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// const creditScoreMin = 500
// const creditScoreMax = 900

type Task struct {
	Id          string `json:"id"`
	Task        string `json:"task"`
	IsCompleted string `json:"iscompleted"`
}
type Result struct {
	Value []Task
	Err   error
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

func taskGet() Result {

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
		if err := rows.Scan(&task.Id, &task.Task, &task.IsCompleted); err != nil {
			fmt.Println("<<<<<<< In the Get Function >>>>>>>>>", err)
			return Result{tasks, err}
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return Result{tasks, err}
	}

	return Result{tasks, nil}

}

func taskPostPut(c *gin.Context, sql string, arg string) {
	var err error
	// Handle the POST request...
	// stmt, err :=
	// if err != nil {
	// 	log.Println(err)
	// 	return false
	//    }
	//    defer stmt.Close()
	// Read the request body
	body, err := io.ReadAll(c.Request.Body)
	// fmt.Println(body)
	println(string(body))
	// if err != nil {
	// 	http.Error(w, "Failed to read request body", http.StatusInternalServerError)
	// 	return
	// }
	defer c.Request.Body.Close()

	// Parse the JSON body
	var task Task

	if err := json.Unmarshal(body, &task); err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
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

// func handler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("<<<<<<< In the handler >>>>>>>>>")
// 	switch r.Method {
// 	case http.MethodGet:
// 		taskGet()
// 	case http.MethodPost:
// 		taskPostPut(w, r, `
// 		INSERT INTO "List" ("task")
// 		VALUES ($1) RETURNING id`, "POST")
// 		fmt.Fprintln(w, "You made a POST request!")
// 	case http.MethodPut:
// 		taskPostPut(w, r, `UPDATE "List" SET "isComplete"=NOT "isComplete","completed_at"=NOW() WHERE "id"=$1 RETURNING id`, "PUT")
// 		fmt.Fprintln(w, "You made a PUT request!")
// 	case http.MethodDelete:
// 		fmt.Fprintln(w, "You made a DELETE request!")
// 		taskPostPut(w, r, ` DELETE FROM "List" WHERE "id"=$1 RETURNING id`, "PUT")
// 	default:
// 		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
// 	}
// }

func main() {

	router := gin.Default()

	router.Use(cors.Default())

	api := router.Group("/")
	{
		api.PUT("/", func(c *gin.Context) {

			taskPostPut(c, `UPDATE "List" SET "isComplete"=NOT "isComplete" WHERE "id"=$1 RETURNING id`, "PUT")

			// taskPostPut(*gin.Context)
			c.JSON(http.StatusOK, nil)
		})
		api.POST("/", func(c *gin.Context) {

			taskPostPut(c, `
			INSERT INTO "List" ("task")
			VALUES ($1) RETURNING id`, "POST")

			// taskPostPut(*gin.Context)
			c.JSON(http.StatusOK, nil)
		})
		api.GET("/", func(c *gin.Context) {

			data := gin.H{
				"message": "Hello from Go backend",
				"status":  "success",
				"data":    taskGet(),
			}
			c.JSON(http.StatusOK, data)
		})
	}

	// Start and run the server
	router.Run(":5000")
	// fmt.Println("New record ID is:", id)
	// http.HandleFunc("/", handler)

	fmt.Println("Server is running on http://localhost:5000")
	// http.ListenAndServe(":8080", nil)
	fmt.Println("Successfully connected!")
}
