package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	var err error
	// Connect to the database
	db, err = sql.Open("postgres", "postgresql://username:password@postgres/db?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// Verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/add", addMessageHandler)
	http.HandleFunc("/read", readMessagesHandler)
	http.HandleFunc("/update", updateMessageHandler)
	log.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func addMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	content := r.FormValue("content")
	if content == "" {
		http.Error(w, "Empty content is not allowed", http.StatusBadRequest)
		return
	}

	_, err := db.Exec("INSERT INTO messages (content) VALUES ($1)", content)
	if err != nil {
		http.Error(w, "Failed to insert message", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message added successfully")
}

func readMessagesHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, content FROM messages")
	if err != nil {
		http.Error(w, "Failed to query messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		if err := rows.Scan(&id, &content); err != nil {
			http.Error(w, "Failed to read message", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "ID: %d, Content: %s\n", id, content)
	}
}

func updateMessageHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse form data
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	// Extract the message ID and new content from the form data
	id := r.FormValue("id")
	content := r.FormValue("content")
	if id == "" || content == "" {
		http.Error(w, "Missing id or content", http.StatusBadRequest)
		return
	}

	// Update the message in the database
	_, err = db.Exec("UPDATE messages SET content = $1 WHERE id = $2", content, id)
	if err != nil {
		http.Error(w, "Failed to update message", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Message updated successfully")
}
