package main

import "fmt"

// Example of a constant block
const (
	// iota valud starts at 0 and increments
	// by 1 each time its used.
	first  = iota + 6
	second = iota
)

// Iota resets once in a new constant block
const (
	third = iota
	fourth
)

func main() {
	fmt.Println(first, second, third, fourth)
}
