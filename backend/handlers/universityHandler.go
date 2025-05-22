package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/qwsnxnjene/kursa4/backend/db"
)

type UniFull struct {
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
	IMG1      string `json:"img1"`
	IMG2      string `json:"img2"`
	Article1  string `json:"article1"`
	Article2  string `json:"article2"`
	URL       string `json:"url"`
}

func UniversityHandler(id string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")

		idInt, _ := strconv.Atoi(id)
		city := r.URL.Query().Get("city")
		var citySql string
		if city == "kazan" {
			citySql = "kazan"
		} else if city == "moscow" {
			citySql = "moscow"
		} else if city == "petersburg" {
			citySql = "stp"
		}

		result := UniFull{}

		query := fmt.Sprintf("SELECT name, short_name, img1, img2, article1, article2, url FROM %s_unis WHERE id = :id", citySql)
		rows, err := db.Database.Query(query, sql.Named("id", idInt))
		if err != nil {
			json.NewEncoder(w).Encode([]Uni{})
			return
		}
		defer rows.Close()

		for rows.Next() {
			var uni UniFull

			err := rows.Scan(&uni.Name, &uni.ShortName,
				&uni.IMG1, &uni.IMG2, &uni.Article1, &uni.Article2, &uni.URL)
			if err != nil {
				json.NewEncoder(w).Encode([]Uni{})
				return
			}
			result = uni
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
