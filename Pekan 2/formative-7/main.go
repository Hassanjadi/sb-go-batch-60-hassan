package main

import "fmt"

func main() {
	// Soal 1
	var buah1 = buah{}
	buah1.nama = "Nanas"
	buah1.warna = "Kuning"
	buah1.adaBijinya = false
	buah1.harga = 9000

	var buah2 = buah{"Jeruk", "Oranye", true, 8000}
	var buah3 = buah{"Semangka", "Hijau & Merah", true, 10000}

	var buah4 = buah{nama: "Semangka",warna: "Hijau & Merah", adaBijinya: true, harga: 10000}

	fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", buah1.nama, buah1.warna, boolToString(buah1.adaBijinya), buah1.harga)
    fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", buah2.nama, buah2.warna, boolToString(buah2.adaBijinya), buah2.harga)
    fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", buah3.nama, buah3.warna, boolToString(buah3.adaBijinya), buah3.harga)
    fmt.Printf("Nama: %s, Warna: %s, Ada Bijinya: %s, Harga: %d\n", buah4.nama, buah4.warna, boolToString(buah4.adaBijinya), buah4.harga)

	// Soal 2
	var s segitiga = segitiga{5, 10}
	fmt.Printf("Luas Segitiga : %d\n", s.luasSegitiga())

	var p persegi = persegi{5}
	fmt.Printf("Luas Persegi : %d\n", p.luasPersegi())

	var pp persegiPanjang = persegiPanjang{5, 3}
	fmt.Printf("Luas Persegi Panjnag : %d\n", pp.luasPersegiPanjang())

	// Soal 3
	myPhone := &phone{
		name:   "Galaxy S21",
		brand:  "Samsung",
		year:   2021,
		colors: []string{"Green Orcid"},
	}

	myPhone.addColor("Black Metalic")

	fmt.Println("Name:", myPhone.name)
	fmt.Println("Brand:", myPhone.brand)
	fmt.Println("Year:", myPhone.year)
	fmt.Println("Color:", myPhone.colors)

	// Soal 4
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, film := range dataFilm {
		fmt.Printf("%d. Title: %s\n   Duration: %d minutes\n   Genre: %s\n   Year: %d\n", i+1, film.title, film.duration, film.genre, film.year)
	}
}

type buah struct {
	nama, warna string
	adaBijinya bool
	harga      uint32
}

func boolToString(adaBijinya bool) string {

	if adaBijinya {
		return "Ada"
	} 
	return "Tidak"
}

type segitiga struct{
  alas, tinggi int
}

func (s segitiga) luasSegitiga() int {
	return s.alas * s.tinggi / 2
}

type persegi struct{
  sisi int
}

func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct{
  panjang, lebar int
}

func (pp persegiPanjang) luasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

type phone struct{
   name, brand string
   year int
   colors []string
}

func (ph *phone) addColor(newColor string) {
	ph.colors = append(ph.colors, newColor)
}

type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, dataFilm *[]movie) {
	newMovie := movie{
		title:    title,
		duration: duration,
		genre:    genre,
		year:     year,
	}

	*dataFilm = append(*dataFilm, newMovie)
}