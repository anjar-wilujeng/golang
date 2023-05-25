package main

import "fmt"

func main() {
	var NilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir = NilaiAkhir > 80
	var lulusabsensi = absensi > 80
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusabsensi)

	var lulus bool = lulusNilaiAkhir && lulusabsensi
	fmt.Println(lulus)

	fmt.Println(NilaiAkhir > 80 && absensi > 80)
}
