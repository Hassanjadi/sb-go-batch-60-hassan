package main

import (
	"fmt"
	. "formative-9/lib"
)

func main() {
	var bangunDatar HitungBangunDatar
	var bangunRuang HitungBangunRuang

	bangunDatar = SegitigaSamaSisi{Alas: 4, Tinggi: 3}
	fmt.Println("===== segitigaSamaSisi")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunDatar = PersegiPanjang{Panjang: 5, Lebar: 3}
	fmt.Println("===== persegiPanjang")
	fmt.Println("luas      :", bangunDatar.Luas())
	fmt.Println("keliling  :", bangunDatar.Keliling())

	bangunRuang = Balok{Panjang: 4, Lebar: 3, Tinggi: 2}
	fmt.Println("===== Balok")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luas permukaan  :", bangunRuang.LuasPermukaan())

	bangunRuang = Tabung{JariJari: 5, Tinggi: 3}
	fmt.Println("===== Tabung")
	fmt.Println("volume      :", bangunRuang.Volume())
	fmt.Println("luas permukaan  :", bangunRuang.LuasPermukaan())

	// Soal 2
	var p PhoneInterface = Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(p.Display())

	// Soal 3
	fmt.Println(LuasPersegi(4, true))
	fmt.Println(LuasPersegi(8, false))
	fmt.Println(LuasPersegi(0, true))
	fmt.Println(LuasPersegi(0, false))

	// Soal 4
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	prefixStr := prefix.(string)
	angkaPertama := kumpulanAngkaPertama.([]int)
	angkaKedua := kumpulanAngkaKedua.([]int)

	total := 0
	output := prefixStr

	for i, angka := range angkaPertama {
		output += fmt.Sprintf("%d", angka)
		total += angka
		if i < len(angkaPertama)-1 {
			output += " + "
		}
	}

	output += " + "

	for i, angka := range angkaKedua {
		output += fmt.Sprintf("%d", angka)
		total += angka
		if i < len(angkaKedua)-1 {
			output += " + "
		}
	}

	output += fmt.Sprintf(" = %d", total)

	fmt.Println(output)
}