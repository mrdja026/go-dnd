package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	dnddata "dnd.go/pkg/service/dnd_data"
	"github.com/supabase-community/supabase-go"
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
	jsonResponse, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	fmt.Print("stignem do ovde")
	w.Write(jsonResponse)
	return

}

func Encouter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encouterId := r.URL.Query().Get("q")
	fmt.Println("GET params were:" + encouterId)
	client, err := supabase.NewClient(API_URL, API_KEY, nil)

	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	data, _, sqlErr := client.From("encounter").Select("*", "GET", false).Eq("id", encouterId).Execute()
	if sqlErr != nil {
		fmt.Println("cannot get encouter", sqlErr)
		panic(sqlErr)
	}
	fmt.Println(data)
	w.WriteHeader(http.StatusFound)
	w.Write([]byte("Encouter returned"))
	return

}
