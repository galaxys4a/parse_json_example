package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

func Ex2() {
	fmt.Println("JSON Arrays")

	birdsJson := `[
		{
		  "species": "pigeon",
		  "description": "likes to perch on rocks"
		},
		{
		  "species":"eagle",
		  "description":"bird of prey"
		}
	]`

	var birds []Bird
	err := json.Unmarshal([]byte(birdsJson), &birds)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Birds : %+v\n", birds)
}
