package json_examples

import (
	"encoding/json"
	"time"

	//"fmt"
	"log"
	"net/http"
)

type MyData struct {
	A int             `json:"a"`
	B string          `json:"b"`
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
		log.Fatal(err.Error())
	}
}

func (m *MyData) MarshalJSON() ([]byte, error) {
	type NewMyData MyData

	result := struct {
		Data      *NewMyData `json:"data"`
		TimeStamp time.Time  `json:"date"`
	}{Data: (*NewMyData)(m), TimeStamp: time.Now()}

	return json.Marshal(result)
}
