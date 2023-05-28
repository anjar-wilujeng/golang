package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32) //contoh melebihi nilai maks
	var nilai8 int8 = int8(nilai32)    //contoh melebihi nilai maks

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	fmt.Println(nilai8)

	var nama = "Junaidi"
	var e byte = nama[0]
	var eString = string(e)

	fmt.Println(nama)
	fmt.Println(eString)

}
