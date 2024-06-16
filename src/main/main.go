package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/mrdja026/go-dnd/api/routes"
)

func main() {
	http.HandleFunc("/", routes.Get_Domain)
	http.HandleFunc("/hello", routes.Get_Hello)

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
