package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/csv"
	"strings"
	"encoding/json"

	"github.com/adrianuf22/learn_golang/arquivos/model"
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
	// @v2 To the end
	// @v3 This file must die!
	nomes.Close() // Close the opened file
	
	// Reading the same file twice
	// Files are streams and once readed everything will go away
	// Trys to read same file now returns nothing
	// Solutions by the moment: Read the file again
	
	// Reading again
	nomesReadedAgain, _ := os.Open("nomes.csv")

	// Wrong - Working with CSV
	nomesCSV := csv.NewReader(nomesReadedAgain)
	conteudo, err := nomesCSV.ReadAll()
	if err != nil {
		fmt.Println("Houve um erro ao ler o conteudo do arquivo CSV. Erro: ", err.Error())
		return
	}

	if 0 == len(conteudo) {
		fmt.Println("O conteúdo está vazio!")
		return
	}

	// Reading line by line
	for indice, linha := range conteudo {
		fmt.Printf("Na linha #%d o conteudo é %s\r\n", indice, linha)
		
		// Reading item by item
		for indiceItem, item := range linha {
			fmt.Printf("No item #%d o valor é %s\r\n", indiceItem, item)
		}
	}

	// @v2
	// nomes.Close()
	// @v3
	nomesReadedAgain.Close()

	// Writing files
	// Read for trird time, so, read again
	nomes, _ = os.Open("nomes.csv")
	nomesCSVToCreatesJSON := csv.NewReader(nomes)
	nomesToCreatesJSON, _ := nomesCSVToCreatesJSON.ReadAll()

	// Creates a file
	arquivoJSON, err := os.Create("nomes.json") // Any file can be create with this method

	// Creates a buffer writer passing the new file to be writted
	jsonWriter := bufio.NewWriter(arquivoJSON)
	// Causes JSON file and an array of names, starts the JSON structure
	jsonWriter.WriteString("[\n") // Add [ and break to next line
	
	// Iterate the CSV file to put data into JSON
	for _, linha := range nomesToCreatesJSON {
		for indiceItem, item := range linha {
			// Nothing awesome, just split an string by delimiter
			nomeSobrenome := strings.Split(item, ";")

			nome := model.Nome{}
			nome.Name = nomeSobrenome[0]
			nome.Lastname =  nomeSobrenome[1]
			nomeJSON, _ := json.Marshal(nome)
			jsonWriter.WriteString(string(nomeJSON))

			// JSON has no trailing commas support, the last must be die
			if indiceItem != (len(linha) - 1) {
				jsonWriter.WriteString(",")
			}
			// Break line
			jsonWriter.WriteString("\r\n")
		}
	}
	
	jsonWriter.WriteString("]") // Closes the JSON Array
	jsonWriter.Flush() // Free buffer to file, or save

	arquivoJSON.Close() // And close the file
}