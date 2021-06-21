package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Bird struct {
	Species     string `json:"birdType"`
	Description string `json:"what it does"`
}

func Ex1() {
	fmt.Println("Encoding structured data")

	bird := &Bird{
		Species: "pigeon",
		Description: "likes to eat seed",
	}

	data, err := json.Marshal(bird)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
