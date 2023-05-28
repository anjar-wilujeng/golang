package main

import "fmt"

func main() {

	var nama1 = "Junaidi"
	var nama2 = "Junaidi"
	var nama3 = "Idris"

	var result bool = nama1 == nama2
	var result2 bool = nama1 == nama3
	fmt.Println(result)
	fmt.Println(result2)

	var value1 = 1
	var value2 = 3

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
