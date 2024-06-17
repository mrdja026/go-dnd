package player

import (
	"fmt"

	"github.com/supabase-community/supabase-go"
)

const API_URL = "https://vpxouvcjjddjciufokhn.supabase.co"
const API_KEY = ""

func Add(name string, race string, class string) error {
	client, err := supabase.NewClient(API_URL, API_KEY, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	resByte, errInt, err := client.From("players").Insert(nil, false, name, race, class).Execute()
	if err != nil {
		fmt.Println("cannot insert player", err)
		panic(err)
	}
	fmt.Println(resByte, errInt)
	return nil
}
