package json

import (
	"fmt"
	"testing"

	"encoding/json"
)

func Benchmark_Parser_Default(b *testing.B) {
	value := `{"str":"test","int":111,"bool":true,"null":null}`
	var result struct {
		Str  string `json:"str"`
		Int  int    `json:"int"`
		Bool bool   `json:"bool"`
		Null any    `json:"null"`
	}
	json.Unmarshal([]byte(value), &result)
	fmt.Println(result.Str)
	fmt.Println(result.Int)
	fmt.Println(result.Bool)
	fmt.Println(result.Null)
}
