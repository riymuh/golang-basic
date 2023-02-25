package main

import "fmt"

func main() {
	resp1, resp2 := getHello()

	fmt.Println(resp1, resp2)
}

func getHello() (name string, age int) {
	name = "riyadh"
	age = 12
	return
}

func getHi(param_age int) (name string, age int) {
	name = "riyadh"
	age = param_age
	return
}
