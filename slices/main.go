package main

import (
	"fmt"
)

func main() {
	// An Slice
	var listaDeNumeros []int
	// Unlike Array, an Slice hasn't size defined and values can be setted at any moment
	// Slices are like Arrays to another languages, as PHP and Javascript
	// Every a new value was setted, all Array is copied to a new Array with a new size 

	// Wrong - Setting a value in a pure Slice
	// listaDeNumeros[0] = 1 // index out of range
	// Right - Setting a value (Or, Appending) in a pure Slice
	listaDeNumeros = append(listaDeNumeros, 1)

	fmt.Println("O que há na lista de números?", listaDeNumeros)

	// Creates an Slice and initialize 3 spaces
	listaDeNumeros = make([]int, 3)
	fmt.Println("O que há na lista de números?", listaDeNumeros)

	// Setting a value in a initialized Slice
	listaDeNumeros[0] = 1
	listaDeNumeros[1] = 2
	listaDeNumeros[2] = 3
	fmt.Println("O que há na lista de números?", listaDeNumeros)

	// Getting the length and capacity of a Slice
	fmt.Println(listaDeNumeros, len(listaDeNumeros), cap(listaDeNumeros))

	// Using For with range to iterate an Slice
	for index, number := range listaDeNumeros {
		fmt.Printf("Quando o index é %d o número é %d\r\n", index, number)
	}
}
