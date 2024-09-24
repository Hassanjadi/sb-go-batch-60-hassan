package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	word1 := "Bootcamp"
	word2 := "Digital"
	word3 := "Skill"
	word4 := "Sanbercode"
	word5 := "Golang"

	result := word1 + " " + word2 + " " + word3 + " " + word4 + " " + word5

	fmt.Println(result)

	// Soal 2
	halo := "Halo Dunia"
	find := "Dunia"
	replace := "Golang"

	newHalo := strings.Replace(halo, find, replace, 2)
	fmt.Println(newHalo)

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var kataKeduaBaru = strings.Replace(kataKedua, "s", "S", 1)
	var kataKetigaBaru = strings.Replace(kataKetiga, "r", "R", 7 )
	var kataKeempatBaru = strings.ToUpper(kataKeempat)

	var semuaKata = kataPertama + " " + kataKeduaBaru + " " + kataKetigaBaru + " " + kataKeempatBaru

	fmt.Println(semuaKata)

	// Soal 4
	var angkaPertama= "8"
	var angkaKedua= "5"
	var angkaKetiga= "6"
	var angkaKeempat = "7"

	var angkaPertamaInteger, _ = strconv.Atoi(angkaPertama)
	var angkaKeduaInteger, _ = strconv.Atoi(angkaKedua)
	var angkaKetigaInteger, _ = strconv.Atoi(angkaKetiga)
	var angkaKeempatInteger, _ = strconv.Atoi(angkaKeempat)

	var jumlah = angkaPertamaInteger + angkaKeduaInteger + angkaKetigaInteger + angkaKeempatInteger
	
	fmt.Println("Jumlah :", jumlah)

	// Soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimatBaru := strings.Replace(kalimat, "halo halo", "Hi Hi", 1)

	hasilKalimat := `"` + kalimatBaru + `"` + " - " + strconv.Itoa(angka)
	
	fmt.Println(hasilKalimat)
}