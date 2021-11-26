package greeting

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HelloText(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	fmt.Fprintf(w, "Hello %s!", name)
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
