package main

import (
	"fmt"

	"github.com/adrianuf22/learn_golang/pacotes/prefixo"
	"github.com/adrianuf22/learn_golang/pacotes/operadora"
)

//NomeDoUsuario Nome do usuario do sistema
var NomeDoUsuario = "Adriano"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Operadora: %s\r\n", operadora.NomeOperadora)
}