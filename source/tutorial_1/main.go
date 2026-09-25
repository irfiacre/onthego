package main

import "fmt"

func main() {
	// Variables
	var intNumber int = 32767 //By default Go uses an int32, but we have access to "int[8,16,32,64]"
	// We also have access to un assigned integers, uint[8...64]
	intNumber = intNumber + 1
	fmt.Println(intNumber)

	var floatNumber float64 = 12345678.9
	fmt.Println("----", floatNumber)

	//  Strings
	var myString string = "Cool String" + "I made!"
	fmt.Println("++++", myString)
}
