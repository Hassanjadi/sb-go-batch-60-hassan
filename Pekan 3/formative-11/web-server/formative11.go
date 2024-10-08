package formative11

import (
	"fmt"
	"math"
	"net/http"
)

func luasLingkaran(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func kelilingLingkaran(r float64) float64 {
	return 2 * math.Pi * r
}

func volumeTabung(r, t float64) float64 {
	return math.Pi * math.Pow(r, 2) * t
}

func Display(w http.ResponseWriter, r *http.Request) {
	var tinggi float64 = 10
	var jariJari float64 = 7

	fmt.Fprintf(w, "Jari-jari = %.f\n", jariJari)
	fmt.Fprintf(w, "Tinggi = %.f\n", tinggi)
	fmt.Fprintf(w, "Luas Lingkaran : %.f\n", luasLingkaran(jariJari))
	fmt.Fprintf(w, "Keliling Lingkaran : %.f\n", kelilingLingkaran(jariJari))
	fmt.Fprintf(w, "Volume Tabung : %.f\n", volumeTabung(jariJari, tinggi))
}