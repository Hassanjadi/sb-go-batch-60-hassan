package lib

import "fmt"

func LuasPersegi(sisi int, tampilkanDetail bool) interface{} {
	if sisi == 0 && tampilkanDetail {
		return "Maaf anda belum menginput sisi dari persegi"
	} else if sisi == 0 && !tampilkanDetail {
		return nil
	}

	luas := sisi * sisi

	if tampilkanDetail {
		return fmt.Sprintf("luas persegi dengan sisi %d cm adalah %d cm", sisi, luas)
	} else {
		return luas
	}
}