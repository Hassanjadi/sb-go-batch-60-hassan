package lib

type Balok struct {
	Panjang, Lebar, Tinggi int
}

func (b Balok) Volume() float64 {
	return float64(b.Panjang) * float64(b.Lebar) * float64(b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return 2 * (float64(b.Panjang)*float64(b.Lebar) +
		float64(b.Panjang)*float64(b.Tinggi) +
		float64(b.Lebar)*float64(b.Tinggi))
}