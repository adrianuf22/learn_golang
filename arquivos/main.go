package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	// To read this file, program must be compiled at the same path of Main
	nomes, err := os.Open("nomes.csv") // Open some file
	if err != nil {
		fmt.Println("Houve um erro ao abrir o arquivo. Erro: ", err.Error())
		return
	}

	scanner := bufio.NewScanner(nomes) // An iterator
	for scanner.Scan() { // Iterate file lines
		fmt.Println("O conteudo da linha é: ", scanner.Text()) // get text from line
	}
 	nomes.Close() // Close the opened file
}