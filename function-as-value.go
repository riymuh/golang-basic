package main

import "fmt"

func main() {
	fmt.Println(getGoodBye("budi"))

	goodBye := getGoodBye
	fmt.Println(goodBye("bu"))
}

func getGoodBye(name string) string {
	return "good bye " + name
}
