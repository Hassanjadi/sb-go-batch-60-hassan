package lib

type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

func (ss SegitigaSamaSisi) Luas() int {
	return ss.Alas * ss.Tinggi / 2
}

func (ss SegitigaSamaSisi) Keliling() int {
	return 3 * ss.Alas
}