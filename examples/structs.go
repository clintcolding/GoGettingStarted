package main

import "fmt"

func main() {
	// This user struct is only available at the main
	// function level

	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u.FirstName)

	u2 := user{ID: 1, FirstName: "Arthur", LastName: "Dent"}
	fmt.Println(u2)
}
