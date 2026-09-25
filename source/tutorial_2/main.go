package main

import (
	"errors"
	"fmt"
)

func main() {
	const newString string = "Hello World"
	newFunction(newString)

	var numb1, numb2 int = 11, 0

	// var intDivRes int = integerDivision(numb1, numb2)
	// fmt.Println("====", intDivRes)

	var result, remainder, err = integerDivisionWithRemainder(numb1, numb2)
	// if err != nil {
	// 	fmt.Printf("Error --- %v", err.Error())
	// 	return
	// } else if remainder == 0 {
	// 	fmt.Printf("--- Remainder is zero")
	// } else {
	// 	fmt.Printf("This is the result %v with remainder %v -", result, remainder)
	// }

	// Switch statements

	switch {
	case err != nil:
		fmt.Printf("Error --- %v", err.Error())
	case remainder == 0:
		fmt.Printf("--- Remainder is zero")
	default:
		fmt.Printf("This is the result %v with remainder %v -", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Printf("--- Remainder is zero")
	case 1, 2:
		fmt.Printf("--- Remainder is one or two")
	default:
		fmt.Printf("Cool")
	}

}

func newFunction(printStmt string) {
	fmt.Println(printStmt)
}

func integerDivision(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func integerDivisionWithRemainder(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Denominator can not be zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator

	return result, remainder, err
}
