package main

import "fmt"

func main() {
	// All items in an array must be the same type

	// Declaring a variable called "arr" that
	// is a 3 element array of integers
	var arr [3]int

	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	fmt.Println(arr)
	fmt.Println(arr[1])

	// Defined with implicit initialization syntax
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr2)
	fmt.Println(arr2[1])
}
