package main

import "fmt"

func main() {
	name := "Junaidi"

	switch name {
	case "Junaidi":
		fmt.Println("Hallo Junaidi")
	case "Idris":
		fmt.Println("Hallo Idris")
	default:
		fmt.Println("Anda Siapa?")
	}

	//switch length := len(name); length > 5 {
	//case true:
	//	fmt.Println("Nama terlalu panjang")
	//case false:
	//	fmt.Println("Nama terlalu pendek")
	//}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama terlalu panjang")
	case length > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")

	}

}
