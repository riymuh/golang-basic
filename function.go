package main

import "fmt"

func main() {
	sayHello()
	sayHi("riyadh", 12)
}

func sayHello() {
	fmt.Println("hahah")
}

func sayHi(name string, age int) {
	fmt.Println(name, " ", age)
}
