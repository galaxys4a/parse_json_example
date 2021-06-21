package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)


type Bird2 struct {
	Species string `json:"birdType"`
	Description string `json:"what it does,omitempty"`
}

func Ex2() {
	fmt.Println("Ignoring Empty Fields")
	
	pigeon := &Bird2{
		Species: "pigeon",
	}

	data, err := json.Marshal(pigeon)
	if err != nil {
		log.Fatal(err)
	}
	
	fmt.Println(string(data))
}
