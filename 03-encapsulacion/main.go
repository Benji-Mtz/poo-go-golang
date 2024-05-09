package main

import "github.com/Benji-Mtz/poo-go-golang/03-encapsulacion/course"

func main() {
	Go := &course.Course{
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
}
