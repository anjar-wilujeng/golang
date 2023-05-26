package main

import "fmt"

func main() {
	var nama [3]string

	nama[0] = "Junaidi"
	nama[1] = "Idris"
	nama[2] = "Dani"

	fmt.Println(nama[0])
	fmt.Println(nama[1])
	fmt.Println(nama[2])

	var values = [3]int{
		10,
		20,
		30,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//mendapatkan panjang array
	fmt.Println(len(nama))
	fmt.Println(len(values))

}
