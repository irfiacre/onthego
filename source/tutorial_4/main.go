// Strings, Runes, & Bytes
package main

import (
	"fmt"
	"strings"
)

func main() {
	var myString string = "résumé"
	var indexed = myString[0]
	fmt.Println("%v, %T", indexed, indexed)
	for idx, val := range myString {
		println("--", idx, val)
	}
	// To make interaction with strings simple, one can use Runes

	var myString2 = []rune("résumé")
	println(myString2)
	// one can declare a rune type using a single quote
	var newRune = `somerune`

	// In go, str
	println(newRune)

	// Building strings
	var stringK = []string{"a", "b", "c"}
	var stringBuilder strings.Builder

	for idx := range stringK {
		stringBuilder.WriteString(stringK[idx])
	}
	var newStr = stringBuilder.String()
	println(newStr)
}
