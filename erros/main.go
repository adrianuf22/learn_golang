package main

import (
	"fmt"
	"encoding/json"

	"github.com/adrianuf22/learn_golang/erros/model"
)

func main() {
	carro := model.Automovel{}
	carro.Name = "GM Camaro"
	carro.Wheels = 5
	carro.Potency = 3.0
	carro.SetValue(80000)

	carroFromJSON, err := json.Marshal(carro)
	if err != nil {
		fmt.Println("[main] Houve um erro na geracao do objeto JSON: ", err.Error())
		return
	}

	fmt.Println("A struct Automovel transformada em JSON: ", string(carroFromJSON)) // Cast to string

	fmt.Println("Definindo preço promocional...")
	if err := carro.SetValue(12000000); err != nil {
		fmt.Println("[main] Houve um erro ao definir o valor do automovel: ", err.Error())
		
		// Checking the type of Error
		if err == model.ErrValorDoAutomovelMuitoAlto {
			fmt.Println("Vendedor, solicite a aprovação ao seu gerente")
		}
		
		return
	}
	fmt.Println("Promoção ativada!")
}
