package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/qwsnxnjene/kursa4/backend/db"
)

func SaveHandler(rw http.ResponseWriter, r *http.Request) {
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

	var user struct {
		Name    string `json:"name"`
		Surname string `json:"surname"`
		Email   string `json:"email"`
	}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(`{"error":"Неверный формат данных"}`))
		return
	}
	user.Email = strings.TrimSpace(user.Email)
	log.Println(user.Email)

	res, err := db.Database.Exec("UPDATE users SET username = :name, surname = :surname WHERE lower(email) = lower(:email)",
		sql.Named("name", user.Name), sql.Named("surname", user.Surname), sql.Named("email", user.Email))
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"Ошибка при обновлении данных"}`))
		return
	}
	log.Println(res.RowsAffected())

	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte(`{"message":"Данные успешно обновлены"}`))
}
