package main

import "fmt"

func main() {
	const firstname string = "riyadh"
	const lastname string = "riyadh"

	// tidak error ketika variable tidak digunakan
	fmt.Println(firstname)

	const (
		age    = 12
		gender = "pria"
	)

	fmt.Println(age)
}
