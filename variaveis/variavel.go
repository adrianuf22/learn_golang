package main

import "fmt"

// A default value is setted by go when is not defined on definition
// @v1 var endereco string // endereco = ""
// @v1 var telefone = "9999-9999"
// @v2 var endereco, telefone string
// @v2 var quantidade int // quantidade = 0
// @v2 var comprou bool // comprou = false
var (
	// Capitalized variable names defines the scope of variable
	// Capitalized: Public
	// Non Capitalized: Private
	// Every public variable must be described (docBlock)
	// @v2 Endereco, telefone 	string // Endereco variable is public and telefone is private
	// Endereco agora esta documentada e é importante
	// Telefone agora é publico, esta documentada e é importante
	Endereco, Telefone 	string
	// @v2 quantidade			int // quantidade = 0
	quantidade, estoque	int // quantidade = 0
	comprou 			bool // comprou = false
)
// Only on 64bits systems
var valor float64 // valor = 0.00
// Unicode characters -> rune type for special characters
var palavras rune

func main() {
	// @v2
	teste := "Valor de teste"

	// @v1 %v can receive any type
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("valor: %0.3f\r\n", valor)
	// @v1 fmt.Printf("telefone: %s\r\n", telefone)
	// @v2
	fmt.Printf("telefone: %s\r\n", Telefone)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	// @v2
	fmt.Printf("palavras: %v\r\n", palavras)
	// @v3
	// At @v2, the variable teste was defined but not used
	// at @v3 the compiler is blocked and returns "teste declared and not used"
	// In GoLang, any variable can be defined and not used
	fmt.Printf("O valor de teste é: %v\r\n", teste)
}
