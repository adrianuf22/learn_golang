package main

import (
	"fmt"

	"github.com/adrianuf22/learn_golang/mapas/model"
)

// Creates a map
var cache map[string]model.Imovel // Unlike variables and struct -
// _ properties, maps value isn't populated by go, at definition 
// the value is null

func main() {
	// To set value, uses make()
	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{
		Name: "Casa pintada",
		Size: 3400,
		Garage: true,
	}
	casa.SetValue(50000)

	// Using the map
	cache["Cache name"] = casa // The name of map is flexible - 
	// _ can be an string, int and an struct too
	// But, for this last case (Structs), performance can be affected
	// Simple names are recomended and more advantaged in perfomance 

	// To read an item from map
	fmt.Println("Minha casa é: ", cache["Cache name"])

	// And to count itens from map, just uses len()
	fmt.Println("Quantos imoveis já foram registrados?", len(cache))

	apto := model.Imovel{
		Name: "Apto novo",
		Size: 4400,
		Garage: false,
	}
	apto.SetValue(95000)

	// Using an struct property as map item name
	cache[apto.Name] = apto
	// To read
	fmt.Println("Minha apartamento é: ", cache[apto.Name])
	
	// Length now
	fmt.Println("Quantos imoveis já foram registrados?", len(cache))

	// Reading all items from map with a loop
	for nomeImovel, imovel := range cache {
		fmt.Printf("Meu imóvel %s é %+v: \r\n", nomeImovel, imovel)
	}

	// Get an item from map also returns a boolean status
	casa, achou := cache["Cache name"] // achou is value
	if achou {
		fmt.Print("A casa foi encontrada!")
	}

	// Removing an item
	if achou {
		delete(cache, "Cache name")
	}
	fmt.Println("Quantos imoveis já foram registrados?", len(cache))
}