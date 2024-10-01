package main

import (
	"fmt"
	"math"
)

func main() {
	// Soal 1
	var luasLingkaran float64
	var kelilingLingkaran float64

	hitungLuasLingkaran(&luasLingkaran, 10)
	hitungKelilingLingkaran(&kelilingLingkaran, 10)

	fmt.Printf("%.2f\n", luasLingkaran)
	fmt.Printf("%.2f\n", kelilingLingkaran)

	// Soal 2
	var sentence string

	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)

	// Soal 3
	var buah = []string{}

	tambahBuah(&buah)

	for i, item := range buah {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	// Soal 4
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. title: %s\n", i+1, film["title"])
		fmt.Printf("   duration: %s\n", film["duration"])
		fmt.Printf("   genre: %s\n", film["genre"])
		fmt.Printf("   year: %s\n\n", film["year"])
	}
}

func hitungLuasLingkaran(luas *float64, r float64) {
	*luas = math.Pi * r * r
}

func hitungKelilingLingkaran (keliling *float64, r float64) {
	*keliling = 2 * math.Pi * r 
}

func introduce (sentence *string, nama string, gender string, pekerjaan string, usia string) {

	var panggilan string

	if gender == "laki-laki" {
		panggilan = "Pak" 
	} else {
		panggilan = "Bu"
	}

	*sentence = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun" 
}

func tambahBuah(buah *[]string) {
	*buah = append(*buah, "Jeruk", "Semangka", "Mangga", "Strawberry", "Durian", "Manggis", "Alpukat")
}

func tambahDataFilm(title string, duration string, genre string, year string, dataFilm *[]map[string]string) {
	film := map[string]string{
		"title":  title,
		"duration": duration,
		"genre":  genre,
		"year":  year,
	}
	*dataFilm = append(*dataFilm, film)
}