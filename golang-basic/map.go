package main

import "fmt"

func main() {
	person := map[string]string{
		"nama":    "Junaidi",
		"address": "Idris",
	}

	person["title"] = "programmer"

	fmt.Println(person)
	fmt.Println(person["nama"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	//membuat map
	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "junaidi idris"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")

	fmt.Println(book)
	fmt.Println(len(book))
}
