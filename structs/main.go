package main

import (
	"fmt"
)

// Defines a new type Imovel
type Imovel struct {
	Size int // Capitalized properties are exported
	Garage bool
	Name string // A default value also is defined when nothing is passed
	value int // And not capitalized is not exported
}

func main() {
	// Creates a variable with the new type
	casa := Imovel{}

	fmt.Printf("%v\r\n", casa)

	
}
