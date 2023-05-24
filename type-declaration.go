package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var nomorktp NoKTP = "111111111"
	var MarriedStatus Married = true
	fmt.Println(nomorktp)
	fmt.Println(MarriedStatus)

}
