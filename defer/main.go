package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	nomes, err := os.Open("nomes.csv")
	// Instead of close at the end, "schedule" now
	defer nomes.Close() // Defer is a tool to helps control the flow
	// When to use and more info:
	// From https://kylewbanks.com/blog/when-to-use-defer-in-go

	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(nomes)
	for scanner.Scan() {
		fmt.Println("O conteudo da linha é: ", scanner.Text())
	}
	// This die and Defer makes happen
	// nomes.Close()
	// ^ This function was added to a queue by Defer and will be executed after returns and before out of function

	// Execution order
	defer fmt.Println("The execution order of Deferred functions is reverse")
	defer fmt.Println("So, I will be printed before of the last print! First, First, First!")
}