package lib

import (
	"fmt"
	"strings"
)

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

func (p Phone) Display() string {
	return fmt.Sprintf("name   : %s\nbrand  : %s\nyear   : %d\ncolors : %s",
		p.Name,
		p.Brand,
		p.Year,
		strings.Join(p.Colors, ", "))
}