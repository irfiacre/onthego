package main

func main() {
	// Arrays, Slices, Maps, Loops
	// var intArr [3]int32 //Array that can hold 3 integers of type int32. Default types here is [0,0,0]

	// fmt.Println(intArr[0], intArr[1:3]) //indexability

	// // Memory of each element
	// fmt.Println(&intArr[0])
	// fmt.Println(&intArr[1])
	// fmt.Println(&intArr[2])

	// // Initializing the array
	// var intArr2 [3]int = [3]int{1, 2, 3} // or use: var intArr2 := [3]int{1, 2, 3}
	// fmt.Println(intArr2)

	// // Slices
	// var mySlice []int32 = []int32{1, 2, -3}
	// fmt.Println(mySlice)

	// mySlice = append(mySlice, 8)
	// fmt.Println(mySlice)

	// var mySlice2 []int32 = []int32{10, 9}
	// mySlice = append(mySlice, mySlice2...)

	// // Making a slice
	// var newSlice []int32 = make([]int32, 2, 3) // (type, len, capacity)
	// fmt.Println(newSlice)

	// // Maps

	// var myMap map[string]int8 = make(map[string]int8)

	// println("---", myMap)

	// // initializing a map
	// var newMap map[string]uint8 = map[string]uint8{"key": 1, "key2": 5}

	// fmt.Println("=====", newMap["key3"])
	// // Maps return a secondary value that is true or false in case a value is/not in the map

	// var _, existInMap = newMap["keyx"]
	// if existInMap {
	// 	println("======")
	// } else {
	// 	println("::::::")
	// }
	// // One can delete a key value pair
	// delete(newMap, "key")

	// println("====", newMap)

	// Looping over maps, slices, arrays

	// // Over Maps
	// for mapKey, mapVal := range newMap {
	// 	println(mapKey, mapVal)
	// }

	// // Over Arrays
	// for idx, val := range intArr2 {
	// 	println(idx, val)
	// }

	// Equivalent to a while loop
	var x int = 0
	for x < 10 {
		println(x)
		x = x + 1
	}
	// or
	for {
		if x >= 10 {
			break
		}
		println(x)

		x = x + 1
	}
}
