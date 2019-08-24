package main

import (
	"fmt"
)

func main() {
	port := 3000
	// Calling the example function w/ args and capturing output
	// An _ represent a write only var
	port, err := startWebServer(port, 2)
	fmt.Println(port, err)
}

// Example function with a parameter and datatype
// int and error represent the data types to be returned
func startWebServer(port int, numberOfRetries int) (int, error) {
	fmt.Println("Starting server...")
	// do important things
	fmt.Println("Server started on port", port)
	fmt.Println("Number of retries", numberOfRetries)
	return port, nil
}
