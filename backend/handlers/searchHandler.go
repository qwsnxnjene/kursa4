package handlers

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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var DB *sql.DB

func OpenDB() *sql.DB {
	var err error
	dbPath := filepath.Join(".", "db", "unis_full.db")
	DB, err = sql.Open("sqlite3", dbPath)

	if err != nil {
		log.Fatal(err)
	}

	return DB
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	query := r.URL.Query().Get("query")
	if query == "" {
		json.NewEncoder(w).Encode([]Uni{})
		return
	}

	result := []Uni{}

	city := r.URL.Query().Get("city")
	query = query + "%"

	if DB == nil {
		fmt.Println("Database connection not initialized")
		return
	}

	if city == "kazan" {
		rows, err := DB.Query("SELECT id, name FROM kazan_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
			sql.Named("query", strings.ToLower(query)), sql.Named("queryStart", "%"+strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.ID, &name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	} else if city == "moscow" {
		rows, err := DB.Query("SELECT id, name FROM moscow_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
			sql.Named("query", strings.ToLower(query)), sql.Named("queryStart", "%"+strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.ID, &name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	} else if city == "petersburg" {
		rows, err := DB.Query("SELECT id, name FROM stp_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
			sql.Named("query", strings.ToLower(query)), sql.Named("queryStart", "%"+strings.ToLower(query)))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			name := Uni{}

			err = rows.Scan(&name.ID, &name.Name)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
			}
			result = append(result, name)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
