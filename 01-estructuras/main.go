package main

import "fmt"

// struct nos servira para declarar el objeto con sus atributos
type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

func main() {
	// Implementamos el llenado de los atriutos en cualquier orden mientras se estipule el atributo-valor
	Go := Course{
		Price:   12.34,
		Name:    "POO en Golang",
		UserIDs: []uint{12, 56, 89},
		IsFree:  false,
		Clases: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Mapas",
		},
	}

	// Si no se estipula el atributo de la clase se debe respetar el orden en que se creo el atributo para su valor
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

	// Si no se especifican todos los atributos tendran su valor cero
	css := Course{
		Name:   "CSS desde cero",
		IsFree: true,
	}

	// Tambien se puede instanciar la clase vacia y hacer un llenado posterior
	js := Course{}

	js.Name = "Curso Javascript"
	js.UserIDs = []uint{12, 1, 2, 3, 4}

	fmt.Println(Go.Name)
	fmt.Println(Go2.IsFree)
	// fmt.Println(css)
	fmt.Printf("%+v\n", css)
	fmt.Printf("%+v\n", js)

	fmt.Println(Go)
}
