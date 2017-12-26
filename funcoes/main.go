package main

import (
	"fmt"

	"github.com/adrianuf22/learn_golang/funcoes/matematica"
)

func main() {
	resultado := matematica.Calc(matematica.Multiplicar, 2, 3)

	fmt.Printf("O resultado de 2 * 3 é igual: %d\r\n", resultado)

	// @v2
	// In this case, uses = instead of :=, cause now we are just defining a value
	// to an created variable
	resultado = matematica.Calc(matematica.Soma, 3, 4)
	fmt.Printf("O resultado de 3 + 4 é igual: %d\r\n", resultado)

	// @v7
	// resultado = matematica.Calc(matematica.Dividir, 2, 2)
	// fmt.Printf("O resultado de 2 / 2 é igual: %d\r\n", resultado)
	// @v8
	// Now, uses := cause resto variable was not defined yet
	resultado, resto := matematica.Dividir(2, 2)
	fmt.Printf("O resultado de 2 / 2 é igual: %d\r\n", resultado)
	fmt.Printf("E o resto de 2 / 2 é: %d\r\n", resto)
}

// @v5 func multiplicar(x int, y int) int {
