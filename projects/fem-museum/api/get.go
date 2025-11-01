package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"frontendmasters.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting the id params here
	id := r.URL.Query()["id"]

	if id != nil {

		// parse the "id" to integer (and catch any ocassional error)
		finalID, err := strconv.Atoi(id[0])

		if err != nil || finalID >= len(data.GetAll()) {

			// error converting the "id" param to integer
			http.Error(w, "Invalid Exhibition Asked", http.StatusBadRequest)
			return

		}

		// finally returning the exhibition for that "id"
		json.NewEncoder(w).Encode(data.GetAll()[finalID])
		return

	}

	// send all if no "id" mentioned
	json.NewEncoder(w).Encode(data.GetAll())

}
