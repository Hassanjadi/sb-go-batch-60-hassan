package main

import "fmt"

func main() {
	// Soal 1
	for i := 1; i <= 20; i++ {
		if i % 2 != 0 && i % 3 == 0 {
			fmt.Println(i, "- I Love Coding")
		} else if i % 2 != 0  {
			fmt.Println(i, "- Santai")
		} else {
			fmt.Println(i, "- Berkualitas")
		}
	}

	// Soal 2
	for i := 1; i <= 7; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("# ")
		}
		fmt.Println()
	}

	// Soal 3
	var kalimat = []string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	var kalimats = kalimat[2:]

	fmt.Println(kalimats)

	// Soal 4
	var sayuran = []string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}


	for i , sayur := range sayuran {
    	fmt.Printf("%d. %s\n", i+1, sayur)
	}

	// Soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

    for key, value := range satuan {
        fmt.Printf("%s = %d\n", key, value)
    }

    volume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
    fmt.Printf("Volume Balok = %d\n", volume)
}