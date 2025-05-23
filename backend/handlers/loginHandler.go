package handlers

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/golang-jwt/jwt"
	"github.com/qwsnxnjene/kursa4/backend/db"
	"golang.org/x/crypto/bcrypt"
)

func authenticateUser(login, password string) (bool, error) {
	var storedHash string
	err := db.Database.QueryRow("SELECT password FROM users WHERE email = ?", login).Scan(&storedHash)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil // Пользователь не найден
	} else if err != nil {
		return false, err // Ошибка базы данных
	}
	err = bcrypt.CompareHashAndPassword([]byte(storedHash), []byte(password))
	if err != nil {
		return false, nil // Пароль не совпадает
	}
	return true, nil
}

func LoginHandler(rw http.ResponseWriter, r *http.Request) {
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
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		rw.Write([]byte(fmt.Sprintf(`{"error":"ошибка десериализации %v"}`, err)))
		return
	}
	defer r.Body.Close()

	success, err := authenticateUser(user.Email, user.Password)
	if err != nil {
		log.Printf("Ошибка базы данных: %v", err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"error":"Внутренняя ошибка сервера"}`))
		return
	}

	var ans string
	if success {
		secret := []byte("secret")

		hash := sha256.Sum256(secret)

		claims := jwt.MapClaims{
			"hash":  hex.EncodeToString(hash[:]),
			"login": user.Email,
		}

		jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		signedToken, err := jwtToken.SignedString(secret)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(fmt.Errorf("failed to sign jwt: %s", err).Error()))
			return
		}

		query := "SELECT username, role, surname FROM users WHERE email = ?"
		row := db.Database.QueryRow(query, user.Email)
		var username, role, surname string
		err = row.Scan(&username, &role, &surname)
		if err != nil {
			if err == sql.ErrNoRows {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write([]byte(`{"error":"Пользователь не найден"}`))
				return
			}
			log.Printf("Ошибка базы данных: %v", err)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(`{"error":"Внутренняя ошибка сервера"}`))
			return
		}

		ans = fmt.Sprintf(`{"token":"%v", "name": "%s", "role": "%s", "surname": "%s", "email": "%s"}`, signedToken, username, role, surname, user.Email)
		rw.WriteHeader(http.StatusOK)
		fmt.Println(signedToken)
	} else {
		rw.WriteHeader(http.StatusBadRequest)
		ans = `{"error":"Неверный пароль или логин"}`
	}
	rw.Write([]byte(ans))
}
