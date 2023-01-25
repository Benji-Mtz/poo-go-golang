package main

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func main() {
	Go := Course{
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

	Go2 := Course{
		"POO en Golang",
		12.34,
		false,
		[]uint{12, 56, 89},
		map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Mapas",
		},
	}

	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	js := Course{}

	js.Name = "Curso Javascript"
	js.UserIDs = []uint{12, 1, 2, 3, 4}

	fmt.Println(Go.Name)
	fmt.Println(Go2.IsFree)
	// fmt.Println(css)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)

	PrintClases(Go)
}
