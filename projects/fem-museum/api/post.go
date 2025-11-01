package api

import (
	"encoding/json"
	"net/http"

	"frontendmasters.com/go/museum/data"
)

func Post(w http.ResponseWriter, r *http.Request) {

	// handle if any other method than POST is called
	if r.Method != http.MethodPost {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
		return
	}
	var exhibition data.Exhibition
	err := json.NewDecoder(r.Body).Decode(&exhibition)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.Add(exhibition)
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))

}
