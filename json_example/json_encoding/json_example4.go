package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

func Ex4() {
	fmt.Println("Encoding maps")

	// The keys need to be strings, the values can be
	// any serializable value
	birdData := map[string]interface{}{
		"birdSounds": map[string]string{
			"pigeon": "coo",
			"eagle":  "squak",
		},
		"total birds": 2,
	}

	data, err := json.Marshal(birdData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
