package main

import (
	"fmt"

	"github.com/Benji-Mtz/poo-go-golang/02-metodos/modelos"
)

func main() {
	Go := &modelos.Course{
		Name:    "POO en Golang",
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Clases: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Mapas",
		},
	}

	Go.PrintClases()
	Go.ChangePrice(300.00)
	fmt.Println(Go.Price)
}
