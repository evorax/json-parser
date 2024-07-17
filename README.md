# json-parser
under development...<br>
this is json parser <br>
## Install <br>
```
go get -u github.com/evorax/json-parser
```
## example code <br>
```go
package main

import (
    "fmt"
    
    json "github.com/evorax/json-parser"
)

func main() {
	value := `{"str":"test","int":111,"bool":true,"null":null}`
	var result struct {
		Str  string `json:"str"`
		Int  int    `json:"int"`
		Bool bool   `json:"bool"`
		Null any    `json:"null"`
	}
	json.ParseJSON(value, &result)
	fmt.Println(result.Str)
	fmt.Println(result.Int)
	fmt.Println(result.Bool)
	fmt.Println(result.Null)
}
```
## Benchmark is <br>
```
Benchmark_Parser_Default-16     1000000000               0.0000224 ns/op               0 B/op          0 allocs/op
Benchmark_Parser-16             1000000000               0.0000196 ns/op               0 B/op          0 allocs/op
```