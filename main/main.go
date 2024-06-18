package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"dnd.go/db"
	"dnd.go/pkg/api/routes"
)

// move this to character
type Player struct {
	name  string
	class string
	race  string
}

func main() {
	http.HandleFunc("/", routes.Get_Domain)
	http.HandleFunc("/hello", routes.Get_Hello)
	http.HandleFunc("/races", corsMiddleware(routes.GetRaces))
	http.HandleFunc("/classes", corsMiddleware(routes.GetClasses))
	http.Handle("/encounter", corsMiddleware(http.HandlerFunc(routes.Encouter)))
	http.HandleFunc("/start", corsMiddleware(validateRequestCharacter(routes.Add)))
	// http.HandleFunc("/createDb", runner)
	http.HandleFunc("/getData", runner)

	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Server started on port 3333\n")
}

// use this for temp endpoint to create db
func runner(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	db.GetData()
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Game started"))

}

// multiple middlewares works
func validateRequestCharacter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func corsMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next(w, r)
	}
}
