package main

import "fmt"

func main() {
	// Constants must be initializaed when declared
	// The constants value must be able to be determined
	// at compile time.
	const pi = 3.1415
	fmt.Println(pi)

	// The compiler will interpret the type as appropriate each
	// time it runs into it.
	const c = 3
	// The compiler is interpreting c as an int
	fmt.Println(c + 3)
	// The compiler is interpreting c as an float32
	fmt.Println(c + 1.2)
}
