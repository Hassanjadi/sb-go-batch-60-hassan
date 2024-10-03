package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Soal 1
	var bangunDatar hitungBangunDatar
	var bangunRuang hitungBangunRuang

	bangunDatar = segitigaSamaSisi{4, 3}
	fmt.Println("===== segitigaSamaSisi")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())
	
	bangunDatar = persegiPanjang{5, 3}
	fmt.Println("===== persegiPanjang")
	fmt.Println("luas      :", bangunDatar.luas())
	fmt.Println("keliling  :", bangunDatar.keliling())

	bangunRuang = balok{4, 3, 2}
	fmt.Println("===== Balok")
	fmt.Println("volume      :", bangunRuang.volume())
	fmt.Println("luas permukaan  :", bangunRuang.luasPermukaan())
	
	bangunRuang = tabung{5, 3}
	fmt.Println("===== Tabung")
	fmt.Println("volume      :", bangunRuang.volume())
	fmt.Println("luas permukaan  :", bangunRuang.luasPermukaan())

	// Soal 2
	var p phoneInterface = phone {
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(p.display())

	// Soal 3
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

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

// Soal 1
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type segitigaSamaSisi struct {
	alas, tinggi int
}

func (ss segitigaSamaSisi) luas() int {
	return ss.alas * ss.tinggi / 2
}

func (ss segitigaSamaSisi) keliling() int {
	return 3 * ss.alas
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

func (pp persegiPanjang) keliling() int {
	return 2 * (pp.panjang + pp.lebar)
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

type tabung struct {
	jariJari, tinggi float64
}

func (t tabung) volume() float64 {
	return math.Pi * t.jariJari *  t.jariJari * t.tinggi
}

func (t tabung) luasPermukaan() float64 {
	return 2 * math.Pi * t.jariJari * (t.jariJari + t.tinggi)
}

type balok struct {
	panjang, lebar, tinggi int
}

func (b balok) volume() float64 {
	return float64(b.panjang) * float64(b.lebar) * float64(b.tinggi)
}

func (b balok) luasPermukaan() float64 {
	return 2 * (float64(b.panjang) * float64(b.lebar) +
                float64(b.panjang) * float64(b.tinggi) +
                float64(b.lebar) * float64(b.tinggi))
}

// Soal 2
type phoneInterface interface {
	display() string
}

type phone struct{
   name, brand string
   year int
   colors []string
}

func (p phone) display() string {
	return fmt.Sprintf("name   : %s\nbrand  : %s\nyear   : %d\ncolors : %s",
		p.name,
		p.brand,
		p.year,
		strings.Join(p.colors, ", "))
}

// Soal 3
func luasPersegi(sisi int, tampilkanDetail bool) interface{} {
	if sisi == 0 && tampilkanDetail {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && !tampilkanDetail {
		return nil
	}

	luas := sisi * sisi

	if tampilkanDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}