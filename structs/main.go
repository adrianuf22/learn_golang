package main

import (
	"fmt"
	"encoding/json"

	"github.com/adrianuf22/learn_golang/structs/model"
)

//Imovel é um type que armazena dados de um imovel
type Imovel struct { // Defines a new type Imovel
	Size int // Capitalized properties are exported
	Garage bool
	Name string // A default value also is defined when nothing is passed
	value int // And not capitalized is not exported
}

func main() {
	// Creates a variable with the new type
	casa := Imovel{}
	fmt.Printf("%v\r\n", casa)

	// Creates a variable passing the values to type initialization
	fazenda := Imovel{2000, true, "Uataragaa", 3500}
	fmt.Printf("%v\r\n", fazenda)

	// Defining a single or partial value to type initialization
	// Wrong
	// apartamento := Imovel{55}
	// fmt.Printf("%v\r\n", apartamento)

	// Right
	// Uses an spaces between type name and braces
	apartamento := Imovel {
		value: 55, // Trailling commas are required and nice!
	}
	fmt.Printf("%v\r\n", apartamento)

	// Passing values in different order
	chacara := Imovel {
		Garage: true,
		value: 400,
		Name: "Mirapora",
		Size: 680,
	}
	fmt.Printf("%v\r\n", chacara)

	// Setting a property value
	casa.Name = "Adams family's house"
	fmt.Printf("%+v\r\n", casa) // +v prints the value and the data structure. key:value defition of the struct

	// @v9 Struct methods
	// Using a method to set value to a private property
	carro := model.Automovel{}
	carro.Name = "GM Camaro"
	carro.Wheels = 5
	carro.Potency = 3.0
	carro.SetValue(80000)

	fmt.Printf("%+v\r\n", carro)

	// Using a method to get a value
	fmt.Printf("O valor do carro é %d\r\n", carro.GetValue())

	carroFromJSON, _ := json.Marshal(carro)

	// Getting values from a JSON tagged struct
	fmt.Printf("A struct Automovel transformada em JSON: %s", carroFromJSON)
}
