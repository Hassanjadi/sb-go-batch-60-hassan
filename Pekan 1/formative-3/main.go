package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	var panjangPersegiPanjangInt, _ =  strconv.Atoi(panjangPersegiPanjang)
	var lebarPersegiPanjangInt, _ = strconv.Atoi(lebarPersegiPanjang)
	var alasSegitigaInt, _ = strconv.Atoi(alasSegitiga)
	var tinggiSegitigaInt, _ = strconv.Atoi(tinggiSegitiga)


	var luasPersegiPanjang int = panjangPersegiPanjangInt * lebarPersegiPanjangInt
	var kelilingPersegiPanjang int = 2 * (panjangPersegiPanjangInt + lebarPersegiPanjangInt)
	var luasSegitiga int  = (alasSegitigaInt * tinggiSegitigaInt) / 2

	fmt.Println("Luas Persegi Panjnag:", luasPersegiPanjang)
	fmt.Println("Keliling Persegi Panjang:", kelilingPersegiPanjang)
	fmt.Println("Luas Segitiga:", luasSegitiga)

	// Soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("Indeks Nilai John: A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("Indeks Nilai John: B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70{
		fmt.Println("Indeks Nilai John: C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60{
		fmt.Println("Indeks Nilai John: D")
	} else {
		fmt.Println("Indeks Nilai John: E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("Indeks Nilai Doe: A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("Indeks Nilai Doe: B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70{
		fmt.Println("Indeks Nilai Doe: C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60{
		fmt.Println("Indeks Nilai Doe: D")
	} else {
		fmt.Println("Indeks Nilai Doe: E")
	}

	// Soal 3
	var tanggal = 25
	var bulan = 5
	var tahun = 2001

	var namaBulan string
	switch bulan {
	case 1 : 
		namaBulan = "Januari"
	case 2 : 
		namaBulan = "Februari"
	case 3 : 
		namaBulan = "Maret"
	case 4 : 
		namaBulan = "April"
	case 5 : 
		namaBulan = "Mei"
	case 6 : 
		namaBulan = "Juni"
	case 7 : 
		namaBulan = "Juli"
	case 8 : 
		namaBulan = "Agustus"
	case 9 : 
		namaBulan = "September"
	case 10 : 
		namaBulan = "Oktober"
	case 11 : 
		namaBulan = "November"
	case 12 :
		namaBulan = "Desember"
	default:
		namaBulan = "Bulan tidak valid"
	}

	fmt.Printf("%d %s %d\n", tanggal, namaBulan, tahun)

	// Soal 4
	switch {
	case tahun >= 1944  && tahun < 1964 : 
		fmt.Println("Baby Boomber")
	case tahun >= 1965  && tahun < 1979 :
		fmt.Println("Generasi X")
	case tahun >= 1980  && tahun < 1994 : 
		fmt.Println("Generasi Y (Millenials)")
	case tahun >= 1995  && tahun < 2015 : 
		fmt.Println("Generasi Z")
	default :
		fmt.Println("Tahun tidak valid")
	}
}
