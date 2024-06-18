package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Player struct {
	Race_id  string `json:"race_id"`
	Class_id string `json:"class_id"`
	Status   string `json:"status"`
	Name     string `json:"name"`
}

func Add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var player Player
	err := json.NewDecoder(r.Body).Decode(&player)
	fmt.Printf("Received player: %+v\n", player)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Game started"))
	return

}
