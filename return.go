package main

import "fmt"

func main() {
	result := getHello("budi")
	fmt.Println(result)
}

func getHello(name string) string {
	return "hello " + name
}
