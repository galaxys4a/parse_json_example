package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Bird struct {
	Species     string
	Description string
}

func Ex1() {
	fmt.Println("Decoding JSON to Structs")

	birdJson := `{
		"species": "pigeon",
		"description": "likes to perch on rocks"
	}`
	
	bird := Bird{}
	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		log.Fatal(err)
	}

	// the plus flag (%+v) adds field names
	fmt.Printf("%+v\n", bird)
}
