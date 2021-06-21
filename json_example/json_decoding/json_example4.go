package json_encoding

import (
	"encoding/json"
	"fmt"
	"log"
)

func Ex4() {
	fmt.Println("Primitives")
	
	numberJson := "3"
	floatJson := "3.1412"
	stringJson := `"bird"`

	var n int
	var pi float64
	var str string
	var err error

	err = json.Unmarshal([]byte(numberJson), &n)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)

	err = json.Unmarshal([]byte(floatJson), &pi)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pi)

	err = json.Unmarshal([]byte(stringJson), &str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}
