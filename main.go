package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	router.Get("/api/v1/greeting", HelloText)

	router.Post("/api/v2/greeting", HelloJSON)

	http.ListenAndServe(":3000", router)
}

func HelloText(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello %s", name)
}

func HelloJSON(w http.ResponseWriter, r *http.Request) {
	name := struct {
		Name string `json:"name,omitempty"`
	}{}
	json.NewDecoder(r.Body).Decode(&name)

	greeting := struct {
		Message string `json:"message,omitempty"`
	}{fmt.Sprintf("Hello %s", name.Name)}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&greeting)
}
