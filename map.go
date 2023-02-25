package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "riyadh",
		"address": "mandiangin",
	}

	person["gender"] = "pria"

	fmt.Println(person)
	fmt.Println(person["name"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "riyadh"

	fmt.Println(book)

	delete(book, "author")

	fmt.Println(book)
}
