package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Institution struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Mock data - в реальном приложении заменить на запрос к БД
var institutions = []Institution{
	{1, "Московский Государственный Университет"},
	{2, "Санкт-Петербургский Государственный Университет"},
	{3, "Казанский Федеральный Университет"},
	{4, "Новосибирский Государственный Университет"},
	{5, "Московский Физико-Технический Институт"},
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	fmt.Print("Oopsie")
	if query == "" {
		json.NewEncoder(w).Encode([]Institution{})
		return
	}

	var results []Institution
	lowerQuery := strings.ToLower(query)

	for _, inst := range institutions {
		lowerName := strings.ToLower(inst.Name)

		// Ищем совпадения (в реальном приложении лучше использовать полнотекстовый поиск)
		if strings.HasPrefix(lowerName, lowerQuery) || strings.Contains(lowerName, lowerQuery) {
			results = append(results, inst)
		}
	}

	// Ограничиваем количество результатов (сортировка делается на клиенте)
	if len(results) > 5 {
		results = results[:5]
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	json.NewEncoder(w).Encode(results)
}

func main() {
	http.HandleFunc("/api/search", searchHandler)
	http.ListenAndServe(":8081", nil)
}
