package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/qwsnxnjene/kursa4/backend/db"
	"golang.org/x/crypto/bcrypt"
)

func SignUpHandler(rw http.ResponseWriter, r *http.Request) {
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

	// Структура для входных данных
	var p struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	// Десериализация JSON
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		log.Printf("Ошибка базы данных: %v", err)
		rw.Write([]byte(fmt.Sprintf(`{"error":"Ошйбка десерйалйзацйй: %v"}`, err)))
		return
	}
	defer r.Body.Close()

	// Валидация данных
	if len(p.Name) < 2 {
		rw.Write([]byte(`{"error":"Имя должно быть не короче 2 сймволов"}`))
		return
	}
	if len(p.Password) < 8 {
		rw.Write([]byte(`{"error":"Пароль должен быть не короче 8 сймволов"}`))
		return
	}

	rows, err := db.Database.Query(`SELECT COUNT(*) FROM users WHERE email = :email`, sql.Named("email", p.Email))
	var count int
	if err != nil {
		log.Printf("Ошибка базы данных: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"Внутренняя ошйбка сервера"}`))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			log.Printf("Ошибка базы данных: %v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(`{"error":"Внутренняя ошйбка сервера"}`))
			return
		}
		if count > 0 {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`{"error":"Пользователь с таким email уже существует"}`))
			return
		}
	}
	// Хэширование пароля
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Ошибка хэширования пароля: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"Внутренняя ошйбка сервера"}`))
		return
	}

	// Вставка пользователя в базу данных
	_, err = db.Database.Exec("INSERT INTO users (username, password, email, role) VALUES (:name, :password, :email, :role)",
		sql.Named("name", p.Name), sql.Named("password", string(hashedPassword)), sql.Named("email", p.Email), sql.Named("role", p.Role))
	if err != nil {
		log.Printf("Ошибка базы данных: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"Внутренняя ошйбка сервера"}`))
		return
	}

	// err = createNewUser(p.Login)
	// if err != nil {
	// 	rw.WriteHeader(http.StatusInternalServerError)
	// 	rw.Write([]byte(`{"error": "Внутренняя ошибка сервера"}`))
	// }

	rw.WriteHeader(http.StatusCreated)
	rw.Write([]byte(`{"message":"Вы успешно зарегйстрйрованы"}`))
}
