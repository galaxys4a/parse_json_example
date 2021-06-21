package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Bird3 struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func Ex5() {
	fmt.Println("Custom attribute names")
	
	birdJson := `{
		"birdType": "pigeon",
		"what it does": "likes to perch on rocks"
	}`
	
	bird := Bird3{}
	err := json.Unmarshal([]byte(birdJson), &bird)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", bird)
}
