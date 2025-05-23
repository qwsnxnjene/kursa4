package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qwsnxnjene/kursa4/backend/db"
)

type Uni struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

	if db.Database == nil {
		fmt.Println("Database connection not initialized")
		return
	}

	if city == "kazan" {
		rows, err := db.Database.Query("SELECT id, name FROM kazan_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
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
		rows, err := db.Database.Query("SELECT id, name FROM moscow_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
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
		rows, err := db.Database.Query("SELECT id, name FROM stp_unis WHERE name LIKE :query OR name LIKE :queryStart LIMIT 3",
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
