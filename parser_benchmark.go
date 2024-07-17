package json

import (
	"fmt"
	"testing"
)

func Benchmark_Parser(b *testing.B) {
	value := `{"str":"test","int":111,"bool":true,"null":null}`
	var result struct {
		Str  string `json:"str"`
		Int  int    `json:"int"`
		Bool bool   `json:"bool"`
		Null any    `json:"null"`
	}
	ParseJSON(value, &result)
	fmt.Println(result.Str)
	fmt.Println(result.Int)
	fmt.Println(result.Bool)
	fmt.Println(result.Null)
}
