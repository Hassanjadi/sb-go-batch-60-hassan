package lib

type PersegiPanjang struct {
	Panjang, Lebar int
}

func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return 2 * (pp.Panjang + pp.Lebar)
}