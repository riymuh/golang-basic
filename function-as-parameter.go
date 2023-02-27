package main

import "fmt"

func main() {
	sayHelloWithFilter("anjing", spamFilter)
	sayHelloWithFilter("budi", spamFilter)
}

type Filter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	nameFiltered := filter(name)
// 	fmt.Println("hello ", nameFiltered)
// }
func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}
