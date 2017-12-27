package main

import (
	"fmt"
	"encoding/json"
)

type Imovel struct {
	Size int
	Garage bool
	Name string
	Value int
}

func main() {
	casa := Imovel{}
	casa.Name = "Casa de Itarangueta"
	casa.Garage = true
	casa.Size = 300
	casa.Value = 80000

	casaFromJSON, _ := json.Marshal(casa)
	fmt.Println("A struct Imovel transformada em JSON: ", string(casaFromJSON))
}
