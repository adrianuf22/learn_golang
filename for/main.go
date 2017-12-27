package main

import (
	"fmt"
)

func main() {
	// In GO, only for is available for loops

	// Simple for
	for i := 0; i <= 5; i++ {
		fmt.Println("Qual é o valor de i? ", i)
	}
	// Like as IF, variables defined to initialize 
	// aren't available for entire scope, just inside the braces
	fmt.Println("Qual é o valor de i agora? ", i)

	// Declaring indice variable outside For initializer
	ix := 1
	for ; ix <= 5; ix++ {
		fmt.Println("Qual é o valor de ix? ", ix)
	}
	// In this case, the variable is available out and in For statement
	fmt.Println("Qual é o valor de ix agora? ", ix)

	// Using just a condition
	ix = 6
	for ix < 10 {
		ix++
		fmt.Println("Qual é o valor de ix? ", ix)
	}

	// While true
	ix = 0
	keep := true
	for keep {
		ix++
		if ix == 5 {
			keep = false
		}
		fmt.Println("Qual é o valor de ix? ", ix)
	}

	// Waiting for a break
	ix = 0
	for {
		ix++
		if ix > 15 {
			break
		}
		fmt.Println("Qual é o valor de ix? ", ix)
	}
}