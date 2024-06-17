package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	dnddata "dnd.go/pkg/service/dnd_data"
)

func Get_Domain(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func Get_Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	resp := dnddata.GetClasses()
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResponse)
	return

}

func GetRaces(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("GET params were:", r.URL.Query())
	resp := dnddata.GetRaces()
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	w.Write(jsonResponse)
	return

}

func GetClasses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("GET params were:", r.URL.Query())
	resp := dnddata.GetClasses()
	time.Sleep(2 * time.Second)
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Print("stignem do ovde")
	w.Write(jsonResponse)
	return

}
