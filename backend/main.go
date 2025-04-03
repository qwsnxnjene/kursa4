package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type Uni struct {
	Name string `json:"name"`
}

var db *sql.DB

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	if query == "" {
		json.NewEncoder(w).Encode([]Uni{})
		return
	}

	result := []Uni{}

	city := r.URL.Query().Get("city")
	query = "%" + query + "%"

	if db == nil {
		fmt.Println("Database connection not initialized")
		return
	}

	if city == "kazan" {
		rows, err := db.Query("SELECT name FROM kazan_unis WHERE name LIKE :query LIMIT 3",
			sql.Named("query", strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	} else if city == "moscow" {
		rows, err := db.Query("SELECT name FROM moscow_unis WHERE name LIKE :query LIMIT 3",
			sql.Named("query", strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	} else if city == "petersburg" {
		rows, err := db.Query("SELECT name FROM stp_unis WHERE name LIKE :query LIMIT 3",
			sql.Named("query", strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	json.NewEncoder(w).Encode(result)
}

func main() {
	var err error
	dbPath := filepath.Join("db", "unis.db")

	db, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	http.HandleFunc("/api/search", searchHandler)
	http.ListenAndServe(":8081", nil)
}
