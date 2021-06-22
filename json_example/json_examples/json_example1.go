package json_examples

import (
	"encoding/json"
	//"fmt"
	"log"
	"net/http"
)

type MyData struct {
	A int `json:"a"`
	B string `json:"b"`
	C json.RawMessage `json:"c"`
}

func Ex1() {
	http.HandleFunc("/print_json", func(w http.ResponseWriter, r *http.Request) {
		myData := new(MyData)
		json.NewDecoder(r.Body).Decode(myData)
		//fmt.Fprint(w, myData)
		json.NewEncoder(w).Encode(myData)
	})

	err := http.ListenAndServe(":8787", nil)
	if err != nil {
		log.Fatal(err)
	}
}
