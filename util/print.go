package util

import (
	"encoding/json"
	"fmt"
)

// Prints a struct as Json, each value on a new line.
// TODO: Add check for input type
func PrintJson(input ...interface{}) {
	b, _ := json.MarshalIndent(input, "", " ")
	fmt.Println(string(b))
}
