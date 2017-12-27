package main

import (
	"fmt"
)

func main() {
	idade := 28

	// Simple switch
	switch idade {
	case 16:
		fmt.Println("Você já pode votar mas é só isso!")
		// break was not necessary in Go
	case 18:
		fmt.Println("Você é um adulto, cuidado com a responsabilidade adquirida!")
	case 21:
		fmt.Println("Você está ficando velho!")
	default:
		fmt.Println("Não tenho nada para você.")
	}

	// Sibling cases
	switch idade {
	case 21:
		// Instead of Break (Whos needs break?) Go has a keyword to sibling results
		fallthrough // Matching with 21, will execute the line 29
	case 30:
		fallthrough // Like your brother 30 here
	case 50:
		fmt.Println("Você está ficando velho!") // And here also
	default:
		fmt.Println("Não tenho nada para você.")
	}

	// Conditional cases
	switch {
	case 16 == idade:
		fmt.Println("Você já pode votar mas é só isso!")
	case 18 <= idade && 21 > idade:
		fmt.Println("Você é um adulto, cuidado com a responsabilidade adquirida!")
	case 21 <= idade && 27 > idade:
		fmt.Println("Ok, já está votando, pode dirigir, deve estar na faculdade, quem diria?")
	case 27 <= idade:
		fmt.Println("Você está ficando velho!")
	default:
		fmt.Println("Não tenho nada para você.")
	}
}