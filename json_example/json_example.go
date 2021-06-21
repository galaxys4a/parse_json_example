package json_example

import (
	"fmt"
	enc "parse_json_example/json_example/json_encoding"
	dec "parse_json_example/json_example/json_decoding"
)

func Ex() {
	fmt.Println("Structured data:")

	dec.Ex1()
	fmt.Println()

	dec.Ex2()
	fmt.Println()

	dec.Ex3()
	fmt.Println()
	
	dec.Ex4()
	fmt.Println()

	dec.Ex5()
	fmt.Println()

	dec.Ex6()
	fmt.Println()

	fmt.Println("Encoding JSON from Go data:")

	enc.Ex1()
	fmt.Println()

	enc.Ex2()
	fmt.Println()

	enc.Ex3()
	fmt.Println()

	enc.Ex4()
	fmt.Println()
}
