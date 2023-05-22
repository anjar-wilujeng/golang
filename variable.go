package main

import "fmt"

func main() {
	var name string

	name = "Peter"
	fmt.Println(name)

	name = "Sondak"
	fmt.Println(name)

	name = "Jadidi"
	fmt.Println(name)

	//tanpa deklarasi tipe variable

	var contohnama = "Budi"
	fmt.Println(contohnama)

	var age = 40
	fmt.Println(age)

	//tanpa deklarasi var

	country := "Indonesia"
	fmt.Println(country)

	country = "Timor"
	fmt.Println(country)

	//deklarasi multiple variable
	var (
		firstname = "Junaidi"
		lastname  = "Idris"
	)

	fmt.Println(firstname)
	fmt.Println(lastname)

}
