package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

type Data map[string]interface{}

type Data2 struct {
	Birds   map[string]string
	Animals string
}

func Ex6() {
	fmt.Println("Decoding JSON to maps")

	birdJson := `{
		"birds": {
		  "pigeon": "likes to perch on rocks",
		  "eagle": "bird of prey"
		},
		"animals": "none"
	}`

	var err error

	data := make(Data)
	err = json.Unmarshal([]byte(birdJson), &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data)

	// The object stored in the "birds" key is also stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := data["birds"].(map[string]interface{})
	for bird, description := range birds {
		fmt.Println(bird, description.(string))
	}
	fmt.Println(data["animals"].(string))

	data2 := Data2{}
	err = json.Unmarshal([]byte(birdJson), &data2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", data2)
}
