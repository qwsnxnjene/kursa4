package main

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/qwsnxnjene/kursa4/backend/handlers"
)

var DB *sql.DB

func main() {
	db := handlers.OpenDB()
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/api/search", handlers.SearchHandler)
	r.HandleFunc("/api/universities/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]
			handlers.UniversityHandler(id)(w, r)
		})
	http.ListenAndServe(":8081", r)
}
