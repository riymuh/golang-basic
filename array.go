package main

import "fmt"

func main() {
	var data [3]string

	data[0] = "riyadh"
	data[1] = "muh"
	data[2] = "haha"

	fmt.Println(data)
	fmt.Println(data[0])

	var values = [3]int{
		1,
		23,
		45,
	}

	var panjangarray = len(values)
	fmt.Println(values, panjangarray)
}
