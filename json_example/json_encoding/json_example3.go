package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)


func Ex3() {
	fmt.Println("Encoding Arrays and Slices")
	
	pigeon := &Bird{
		Species:     "Pigeon",
		Description: "likes to eat seed",
	}

	pigeons := []*Bird{pigeon, pigeon}

	data, err := json.Marshal(pigeons)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
