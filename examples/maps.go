package main

import "fmt"

func main() {
	// Creating a map with a string data type as the key
	// and an integer data type value
	m := map[string]int{"foo": 42}
	fmt.Println(m)
	fmt.Println(m["foo"])

	m["foo"] = 27
	fmt.Println(m)

	delete(m, "foo")
	fmt.Println(m)
}
