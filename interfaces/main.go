package main

import (
	"fmt"

	"github.com/adrianuf22/learn_golang/interfaces/model"
)

func main() {
	galinhaPata := model.Ave{
		Name: "Ornita",
	}

	galinhaPata.Voa()

	fmt.Println(galinhaPata)
	// cacarejar(galinhaPata)
	// @v2 Interfaces and pointers
	// from https://medium.com/@agileseeker/go-interfaces-pointers-4d1d98d5c9c6
	cacarejar(&galinhaPata) // Now, I need pointer of Galinha, cause one of methods needs a pointer and not just a value
	definirMilhoComoTipoDeComida(&galinhaPata)
	fmt.Println(galinhaPata)

	grasnar(galinhaPata)
}

// Receive an interface Galinha
func cacarejar(galinha model.Galinha) {
	// And expect for method Cacareja is implemented
	fmt.Println(galinha.Cacareja())
}

func grasnar(pata model.Pata) {
	fmt.Println(pata.Grasna())
}

func definirMilhoComoTipoDeComida(galinha model.Galinha) {
	galinha.TipoComida("Milho")
}

