package main

import "fmt"

func main() {
	var bulan = [...]string{
		"januari",
		"februari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice1 = bulan[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	bulan[5] = "juni berubah"
	fmt.Println(slice1)

	slice1[2] = "julii berubah"
	fmt.Println(bulan)

	var slice2 = bulan[10:]

	fmt.Println(slice2)
	fmt.Println(cap(slice2))

	var slice3 = append(slice2, "riyadh")

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
}
