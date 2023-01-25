package modelos

import "fmt"

type Course struct {
	Name    string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Clases  map[uint]string
}

// Metodo de Course c-argumento receptor
func (c Course) PrintClases() {
	text := "Las clases son: "
	for _, class := range c.Clases {
		text += class + ", "
	}
	fmt.Println(text[:len(text)-2])
}
