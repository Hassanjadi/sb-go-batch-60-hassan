package main

import (
	"fmt"
	"strings"
)


func main() {
	// Soal 1
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// Soal 2
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john)

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah)

	// Soal 3
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// Soal 4
	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(genre string, jam string, tahun string, title string) {
		var mapFilm = map[string]string{
			"genre": genre,
			"jam": jam,
			"tahun": tahun,
			"title": title,
		}

		dataFilm = append(dataFilm, mapFilm)
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
	fmt.Println(item)
	}
}


func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}

func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

func introduce (nama, gender, pekerjaan, usia string) (kalimat string) {
	var panggilan string

	if gender == "laki-laki" {
		panggilan = "Pak" 
	} else {
		panggilan = "Bu"
	}

	kalimat = panggilan + " " + nama + " adalah seorang " + pekerjaan + " yang berusia " + usia + " tahun" 
  	return kalimat
}

func buahFavorit(nama string, buah ...string) (kalimat string) {
	var jumlahBuah string

	for i := 0; i < len(buah); i++ {
		if i == len(buah) - 1 {
			jumlahBuah += `"` + buah[i] + `"`
		} else {
			jumlahBuah += `"` + buah[i] + `", `
		}
	}

	kalimat = "halo saya " + strings.ToLower(nama) + " dan buah favorit saya adalah " + jumlahBuah
	return kalimat
}