package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/supabase-community/supabase-go"
)

const API_URL = "https://vpxouvcjjddjciufokhn.supabase.co"
const API_KEY = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InZweG91dmNqamRkamNpdWZva2huIiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTg2Mzc2NTgsImV4cCI6MjAzNDIxMzY1OH0.dzRXzl6hWwJSjQAXKtISoYBJDvuU59C7_rFUPsdozEc"

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
	client, err := supabase.NewClient(API_URL, API_KEY, nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	_, _, sqlErr := client.From("player").Insert(player, false, "", "representation", "exact").Execute()
	if sqlErr != nil {
		fmt.Println("cannot insert player1111", err)
		panic(err)
	}
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Game started"))
	return

}
