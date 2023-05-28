package main

import "fmt"

func main() {
	var nama = "Junaidi"

	if nama == "Junaidi" {
		fmt.Println("Hello Junaidi")
	} else if nama == "Idris" {
		fmt.Println("Idris")
	} else {
		fmt.Println("Hi, Anda siapa?")
	}

	if length := len(nama); length > 5 {
		fmt.Println("Nama sudah benar")
	} else {
		fmt.Println("Nama anda salah")
	}
}
