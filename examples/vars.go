package main

import "fmt"

func main() {
	var i int
	i = 42
	fmt.Println(i)

	var f float32 = 3.14
	fmt.Println(f)

	// Implicit initialization syntax; :=
	// Go compiler will imply the datatype based on value
	firstName := "Arthur"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Println(c)

	r, im := real(c), imag(c)
	fmt.Print(r, im)
}
