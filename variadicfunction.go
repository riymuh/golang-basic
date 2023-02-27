package main

import "fmt"

func main() {
	total := sumAll(10, 20, 10, 30, 10)

	fmt.Println(total)

	slice := []int{10, 20, 20, 40, 40}
	total = sumAll(slice...)

	fmt.Println(total)
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}
