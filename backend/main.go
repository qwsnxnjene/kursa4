package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/qwsnxnjene/kursa4/backend/db"
	"github.com/qwsnxnjene/kursa4/backend/handlers"
)

func main() {
	_, err := db.OpenSql()
	if err != nil {
		log.Fatalf("[main]: %v", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/search", handlers.SearchHandler)
	r.HandleFunc("/api/universities/{id}",
		func(w http.ResponseWriter, r *http.Request) {
			vars := mux.Vars(r)
			id := vars["id"]
			handlers.UniversityHandler(id)(w, r)
		})
	r.HandleFunc("/api/signin", handlers.SignUpHandler)
	r.HandleFunc("/api/login", handlers.LoginHandler)
	r.HandleFunc("/api/save", handlers.SaveHandler)
	r.HandleFunc("/api/adduniversity", handlers.AddUniversityHandler)
	http.ListenAndServe(":8081", r)
}
