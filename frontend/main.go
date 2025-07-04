package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	backendURL := os.Getenv("BACKEND_URL")
	if backendURL == "" {
		log.Fatal("BACKEND_URL must be set")
	}

	log.Printf("Back url: `%s`\n", backendURL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		resp, err := http.Get(backendURL + "/hello")
		if err != nil {
			http.Error(w, "error calling backend", http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "error reading backend response", http.StatusInternalServerError)
			return
		}

		// Simple HTML page showing the message
		fmt.Fprintf(w, `<!DOCTYPE html>
		<html>
			<head>
				<meta charset="UTF-8">
				<title>Hello-World Module</title>
				<style>
				/* Force white background */
				body {
					background-color: white;
					margin: 0;
					padding: 20px;
					font-family: sans-serif;
				}
				</style>
			</head>
			<body>
				<h1>%s</h1>
			</body>
		</html>`, body)
	})

	log.Println("frontend listening on :3000")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
