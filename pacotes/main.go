package main

import (
	"fmt"

	// Importing a local package from GOPATH
	"pacotes/prefixo"
	// Importing a remote package from my repo
	"github.com/adrianuf22/learn_golang/pacotes/operadora"
)

//NomeDoUsuario Nome do usuario do sistema
var NomeDoUsuario = "Adriano"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
}