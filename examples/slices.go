package main

import "fmt"

func main() {
	// Slices are built on top of arrays
	arr := [3]int{1, 2, 3}

	// The colon operator creates a slice from
	// the beginning of the array to the end
	slice := arr[:]

	// Any changes made to the arrary are reflected
	// in the slice and vice versus
	arr[1] = 42
	slice[2] = 27

	fmt.Println(arr, slice)

	// Initializing a slice without an array
	// The compiler will handle creating the
	// underlying array
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)

	// A slice is not a fixed size
	slice2 = append(slice2, 4, 42, 27)
	fmt.Println(slice2)

	// You can create a slice of a slice
	// Starting at index 1 up to the end
	slice3 := slice2[1:]
	// Starting at the beginning up to but not including index 2
	slice4 := slice2[:2]
	// Starting at index 1 up to but not including index 2
	slice5 := slice2[1:2]
	fmt.Println(slice3, slice4, slice5)
}
