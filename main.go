package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

// const creditScoreMin = 500
// const creditScoreMax = 900

type Task struct {
	Task string `json: "task"`
}

const (
	host     = "localhost"
	port     = 5400
	user     = "postgres"
	password = "postgres"
	dbname   = "Todo"
)

// http package gives acess to rest commands
func getCreditScore(w http.ResponseWriter, r *http.Request) {
	var task Task
	w.Write([]byte("new Note"))
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Note: %+v", task)
	// var creditRating = credit_rating{
	// 	// generates a random number between 0 and 400 and adds the min credit score to it
	// 	//creating a json object with credit rating as the key and the random number as the value
	// 	CreditRating: (rand.Intn(creditScoreMax-creditScoreMin) + creditScoreMin),
	// }

	// w.WriteHeader(http.StatusOK)
	// json.NewEncoder(w).Encode(creditRating)
}

func indexGet(w http.ResponseWriter, r *http.Request) {
	// Handle the GET request...
}

func indexPost(w http.ResponseWriter, r *http.Request) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Handle the POST request...
	// Ensure the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

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
	sqlStatement := `
	INSERT INTO "List" ("task")
	VALUES ($1)
	RETURNING id`
	id := 0
	err = db.QueryRow(sqlStatement, task.Task).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:", id)
	// Process the data (for demonstration, just print it)
	fmt.Printf("Received: %+v\n", task)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	response := map[string]string{
		"status":  "success",
		"message": fmt.Sprintf("Hello, %s!", task.Task),
	}
	json.NewEncoder(w).Encode(response)

}

func handleRequests() {
	router := httprouter.New()

	router.HandlerFunc("GET", "/", indexGet)
	router.HandlerFunc("POST", "/todo", indexPost)
	http.Handle(" /todo", http.HandlerFunc(getCreditScore))
	log.Fatal(http.ListenAndServe(":8080", router))
	// http.HandleFunc("POST /todo", getCreditScore)
	// http.HandleFunc("GET /", getCreditScore)
}

func main() {

	// fmt.Println("New record ID is:", id)
	fmt.Println("Successfully connected!")
	handleRequests()
}
