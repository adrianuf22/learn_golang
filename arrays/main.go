package main

import (
	"fmt"
)

func main() {
	// Fixed capacity
	// = can be omitted on definition
	var listaDeNumeros [3]int

	listaDeNumeros[0] = 1
	listaDeNumeros[1] = 2
	listaDeNumeros[2] = 3

	fmt.Println("O que há na lista de números?", listaDeNumeros)

	// Using len (again) to read the size of array
	fmt.Println("Número de itens da lista?", len(listaDeNumeros))

	// Syntax sugar to array create/value definition
	listaDeNomes := [2]string{"Adriano","AZTM"}
	fmt.Println("O que há na lista de nomes?", listaDeNomes)

	// Defining array size by values
	// For 3 values defined, the size will be 3 (really)
	listaComInumerosNome := [...]string{"Adriano", "XPTO", "ABGF"}
	fmt.Println("O que há na lista de varios nomes?", listaComInumerosNome)

	fmt.Println("Número de itens da lista?", len(listaComInumerosNome))

	// Using for with range to iterate an Array
	for index, name := range listaComInumerosNome {
		fmt.Printf("Quando o indice é %d o nome é %+v\r\n", index, name);
	}
}
