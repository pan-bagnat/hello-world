package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// Read connection string from env, or use sensible default
	conn := os.Getenv("DB_CONN")
	if conn == "" {
		conn = "postgres://postgres:postgres@hello-world-db:5432/postgres?sslmode=disable"
	}

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		var msg string
		if err := db.QueryRow("SELECT msg FROM hello LIMIT 1").Scan(&msg); err != nil {
			http.Error(w, "db error", http.StatusInternalServerError)
			return
		}
		w.Write([]byte(msg))
	})

	log.Println("backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
