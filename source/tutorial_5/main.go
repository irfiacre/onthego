// Structures and interfaces

package main

// import (
// 	"fmt"
// 	"strings"
// )

type gasEngine struct {
	kmpl    uint8
	litters uint8
	owner   ownerType
}
type ownerType struct {
	fname string
}

func main() {
	var newEngine gasEngine = gasEngine{10, 5, ownerType{"Adam"}}
	// Anonymous structures
	var myEngine = struct {
		kmpl   uint8
		liters uint8
	}{1, 2}
	newEngine.litters = 20
	println("===", myEngine.liters)

}
