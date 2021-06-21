package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Dimensions struct {
	Height int
	Width  int
}

type Bird2 struct {
	Species     string
	Description string
	Dimensions  Dimensions
}

func Ex3() {
	fmt.Println("Embedded objects")

	birdJson := `{
		"species": "pigeon",
		"description": "likes to perch on rocks",
		"dimensions": {
		  "height": 24,
		  "width": 10
		}
	}`

	bird := Bird2{}
	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", bird)
}
