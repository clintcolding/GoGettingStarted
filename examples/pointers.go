package main

import "fmt"

func main() {
	// Setting a var pointer with type string
	var firstName *string = new(string)

	// Dereferencing the var
	*firstName = "Arthur"
	fmt.Println(firstName, *firstName)

	// Creating a pointer from an existing var
	lastName := "Milton"
	fmt.Println(lastName)

	// Using address of operator; &
	ptr := &lastName
	fmt.Println(ptr, *ptr)

	// When updating the value, the value changes but
	// the memory address remains the same.
	lastName = "Bradley"
	fmt.Println(ptr, *ptr)
}
