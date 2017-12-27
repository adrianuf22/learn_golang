package main

import (
	"fmt"
)

type Imovel struct {
	Size int
	Name string
	Garage bool
	Value int
}

func main() {
	// Creates a reference to memory point with Imovel type
	casa := new(Imovel)
	fmt.Printf("Imovel é: %v\r\n", casa);
	// To view the memory point allocate, uses an & before the variable
	fmt.Printf("Imovel é: %p - %v\r\n", &casa,casa);

	casaNova := Imovel {
		Size: 4000,
		Garage: true,
		Name:"Casa na praia",
	}
	// This method expected a memory point of the type
	// If & is not present, the compiler rejects the struct Imovel
	mudaTamanhoImovel(&casaNova);
	fmt.Printf("Imovel é: %p - %v\r\n", &casaNova,casaNova);
}

// @v8 - implicit assignment of unexported field 'value' in types.Imovel literal
// func mudaValorImovel(imovel *Imovel) {
	// Not exported / private
// 	imovel.value = imovel.value + 50
// }
func mudaTamanhoImovel(imovel *Imovel) { // As reference
	imovel.Size = imovel.Size + 50
}
