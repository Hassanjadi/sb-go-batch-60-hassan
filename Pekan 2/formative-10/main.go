package main

import (
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
)

func main() {
	// Soal 1
	defer susunKalimat("Golang Backend Development", 2021)
    fmt.Println("Fungsi utama sedang dijalankan")

	// Soal 2
	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Soal 3
	angka := 1

	defer cetakAngka(&angka)

	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)

	// Soal 4
	var phones = []string{}

	tambahPhone(&phones, "Xiaomi")
	tambahPhone(&phones, "Asus")
	tambahPhone(&phones, "IPhone")
	tambahPhone(&phones, "Samsung")
	tambahPhone(&phones, "Oppo")
	tambahPhone(&phones, "Realme")
	tambahPhone(&phones, "Vivo")

	tampilkanPhones(&phones)

	// Soal 5
	var phoness = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	var wg sync.WaitGroup

	wg.Add(1)
	go tampilkanPhoness(&wg, phoness)

	wg.Wait()

	// Soal 6
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}

func susunKalimat(kalimat string, tahun int) {
	fmt.Printf("%s, Tahun: %d\n", kalimat, tahun)
}

func kelilingSegitigaSamaSisi(sisi int, tampilkanKalimat bool) (string, error) {
	if sisi == 0 {
		err := errors.New("maaf anda belum menginput sisi dari segitiga sama sisi")
		if !tampilkanKalimat {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Panic terjadi:", r)
				}
			}()
			panic(err)
		}
		return "", err
	}

	if tampilkanKalimat {
		sisi = 2
		keliling := 3 * sisi
		return fmt.Sprintf("Keliling segitiga sama sisinya dengan sisi %d cm adalah %d cm", sisi, keliling), nil
	}

	keliling := 3 * sisi
	return fmt.Sprintf("%d", keliling), nil
}

func tambahAngka(nilai int, angka *int) {
	*angka += nilai
}

func cetakAngka(angka *int) {
	fmt.Println("Total angka:", *angka)
}

func tambahPhone(phones *[]string, phone string) {
	*phones = append(*phones, phone)
}

func tampilkanPhones(phones *[]string) {
	sort.Strings(*phones)

	for i, phone := range *phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func tampilkanPhoness(wg *sync.WaitGroup, phoness []string) {
	defer wg.Done()

	sort.Strings(phoness)

	for i, phone := range phoness {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}
}

func getMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}

	close(moviesChannel)
}