package main

import (
	"fmt"

	"github.com/adrianuf22/learn_golang/funcoes_1/matematica"
)

func main() {
	resultadoDaConta := matematica.Calc(matematica.Multiplicar, 2, 3)

	fmt.Printf("O resultado de 2 * 3 é igual: %d\r\n", resultadoDaConta)

	// @v2
	// In this case, uses = instead of :=, cause now we are just defining a value
	// to an created variable
	resultadoDaConta = matematica.Calc(matematica.Soma, 3, 4)
	fmt.Printf("O resultado de 3 + 4 é igual: %d\r\n", resultadoDaConta)
}

// @v5 func multiplicar(x int, y int) int {
