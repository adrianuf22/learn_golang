package main

import (
	"fmt"

	"pacotes/prefixo"
)

//NomeDoUsuario Nome do usuario do sistema
var NomeDoUsuario = "Adriano"

func main() {
	fmt.Printf("Nome do usuario: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
}