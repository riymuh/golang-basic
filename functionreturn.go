package main

import "fmt"

func main() {
	var result = getHello("budi")
	fmt.Println(result)
}

func getHello(name string) string {
	return "hello " + name
}
