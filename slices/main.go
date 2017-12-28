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

	// @v2
	listaDeNomes := make([]string, 5)
	listaDeNomes[0] = "XPTO"
	listaDeNomes[1] = "ABCF"
	listaDeNomes[2] = "WERT"
	listaDeNomes[3] = "TYPO"
	listaDeNomes[4] = "AURA"

	// Creating an slice of an slice
	// variable[index starts in 0:index starts in 1] <- Rules of GO
	var primeiroNomeDaLista = listaDeNomes[:1]
	fmt.Println("Qual é o nome?", primeiroNomeDaLista)

	var ultimoNomeDaLista = listaDeNomes[len(listaDeNomes) - 1:]
	fmt.Println("Qual é o nome?", ultimoNomeDaLista)

	var nomeDoMeioDaLista = listaDeNomes[2:3]
	fmt.Println("Qual é o nome?", nomeDoMeioDaLista)

	// Removing a item from an Slice (Hell)
	indiceDoNomeASerRemovido := 3 // TYPO
	tmp := listaDeNomes[:indiceDoNomeASerRemovido]
	fmt.Println("O que sobrou?", tmp)
	tmp = append(tmp, listaDeNomes[indiceDoNomeASerRemovido+1:]...)
	fmt.Println("O que sobrou?", tmp)
	copy(listaDeNomes, tmp)
	fmt.Println("Como ficou?", listaDeNomes)
	// Has a problem here - The deleted data is replaced by the next item - Searching fixes

	// From https://github.com/golang/go/wiki/SliceTricks
	// @v2 Removing a item from an Slice
	indiceDoNomeASerRemovido = 3
	tmp = listaDeNomes[:indiceDoNomeASerRemovido]
	fmt.Println("O que sobrou?", tmp)
	listaDeNomes = append(tmp, listaDeNomes[indiceDoNomeASerRemovido+1:]...)
	fmt.Println("Como ficou?", listaDeNomes)

	// With better names
	ixToDelete := 3
	listaDeNomes = append(listaDeNomes[:ixToDelete], listaDeNomes[ixToDelete+1:]...)
	fmt.Println("Como ficou?", listaDeNomes)

	// Cuting items from a Slice
	listaDeNomes = append(listaDeNomes[:1], listaDeNomes[0:])
	fmt.Println("Como ficou?", listaDeNomes)
}
