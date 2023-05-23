package main

import "fmt"

func main() {
	//const namadepan string = "Junadidi"
	//const namabelakang = "Idris"
	//const umur = 40

	// fmt.Println(namadepan)
	//fmt.Println(namabelakang)
	//fmt.Println(umur)

	// multiple constant
	const (
		namadepan    = "Fahmi"
		namabelakang = "idris"
		age          = 40
	)

	fmt.Println(namadepan)
	fmt.Println(namabelakang)
	fmt.Println(age)

	//constant tidak bisa dirubah nilainya

}
