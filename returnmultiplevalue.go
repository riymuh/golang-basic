package main

import "fmt"

func main() {
	resp1, resp2 := getHello("budi", 12)
	resp3, _ := getHello("budi", 12)
	fmt.Println(resp1, resp2)
	fmt.Println(resp3)
}

func getHello(name string, age int) (string, int) {
	return "hello " + name, age
}
