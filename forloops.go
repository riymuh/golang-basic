package main

import "fmt"

func main() {
	counter := 5

	for i := 0; i < counter; i++ {
		fmt.Println(i)
	}

	for counter <= 10 {
		fmt.Println(counter)
		counter++
	}

	//for range

	names := []string{"riyadh", "budi", "adi"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}
