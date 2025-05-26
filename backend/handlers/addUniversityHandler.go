package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/qwsnxnjene/kursa4/backend/db"
)

func AddUniversityHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	rw.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		rw.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"error":"Неверный метод запроса"}`))
		return
	}
	rw.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var university struct {
		ShortName string `json:"short_name"`
		FullName  string `json:"full_name"`
		Img1      string `json:"img1"`
		Img2      string `json:"img2"`
		Article1  string `json:"article1"`
		Article2  string `json:"article2"`
		City      string `json:"city"`
		Url       string `json:"url"`
	}

	err := json.NewDecoder(r.Body).Decode(&university)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf(`{"error":"ошибка десериализации %v"}`, err)))
		return
	}
	defer r.Body.Close()

	city := ""
	if university.City == "kazan" {
		city = "kazan"
	} else if university.City == "moscow" {
		city = "moscow"
	} else if university.City == "petersburg" {
		city = "stp"
	}

	query := fmt.Sprintf(`INSERT INTO %s_unis (short_name, name, img1, img2, article1, article2, url) VALUES (:short, :name, :img1, :img2, :art1, :art2, :url)`, city)
	_, err = db.Database.Exec(query, sql.Named("short", university.ShortName),
		sql.Named("name", university.FullName),
		sql.Named("img1", university.Img1),
		sql.Named("img2", university.Img2),
		sql.Named("art1", university.Article1),
		sql.Named("art2", university.Article2),
		sql.Named("url", university.Url))
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf(`{"error":"ошибка добавления %v"}`, err)))
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"status":"Университет успешно добавлен"}`))
}
