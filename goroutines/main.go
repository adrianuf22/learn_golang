package main

import (
	"fmt"
	"strings"
	"os"
	"bufio"
	"encoding/csv"
	"encoding/json"
	"sync"
	"time"

	"github.com/adrianuf22/learn_golang/goroutines/model"
)

var osquestrador sync.WaitGroup // Before an WaitGroup, any stdout can be getted cause, Go dont wait for concurrents and finish the application

func main() {
	// Register two concurrents to wait
	osquestrador.Add(2)
	// go creates a concurrent
	go converteNomesCSVToJSON("nomes")
	// and now, another concurrent
	go converteNomesCSVToJSON("outros-nomes")
	// Wait for the Done sign
	osquestrador.Wait()

	// Using Channels
	var canal chan string // Creates a channel of strings
	canal = make(chan string) // Initialize the channel

	go ping(canal)
	go pong(canal)
	go printPingPong(canal)

	// Different of Sync orquestrador, this will help to break a goroutine when some enter key was typed during execution
	var entrada string
	fmt.Scanln(&entrada)
}

func converteNomesCSVToJSON(arquivoDeNomes string) {
	fmt.Println("Convertendo o arquivo", arquivoDeNomes, " para JSON...")
	arquivo, err := os.Open(arquivoDeNomes + ".csv")
	defer arquivo.Close()

	if err != nil {
		fmt.Println("Houve um erro ao ler o arquivo CSV. Erro: ", err.Error())
		return
	}

	// Wrong - Working with CSV
	arquivoCSV := csv.NewReader(arquivo)
	conteudo, err := arquivoCSV.ReadAll()
	if err != nil {
		fmt.Println("Houve um erro ao ler o conteudo do arquivo CSV. Erro: ", err.Error())
		return
	}

	if 0 == len(conteudo) {
		fmt.Println("O conteúdo está vazio!")
		return
	}

	arquivoJSON, err := os.Create(arquivoDeNomes + ".json")
	defer arquivoJSON.Close()
	
	jsonWriter := bufio.NewWriter(arquivoJSON)
	defer jsonWriter.Flush()

	jsonWriter.WriteString("[\n")
	
	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			nomeSobrenome := strings.Split(item, ";")

			nome := model.Nome{}
			nome.Name = nomeSobrenome[0]
			nome.Lastname =  nomeSobrenome[1]
			nomeJSON, _ := json.Marshal(nome)
			jsonWriter.WriteString(string(nomeJSON))

			if indiceItem != (len(linha) - 1) {
				jsonWriter.WriteString(",")
			}
			jsonWriter.WriteString("\r\n")
		}
	}
	jsonWriter.WriteString("]")

	fmt.Println("Arquivo criado com sucesso!!!")
	
	// Here, has no more work, so Done
	osquestrador.Done()
	// ^ If Done wasn't called: all goroutines are asleep - deadlock!
}

func ping(canal chan string) {
	for {
		canal <- "ping" + time.Now().String() // With channel, the sign <- represents the direction of data comes from to
			// ^ If is hard to understand, is like as = in Channel
	}
}

func pong(canal chan string) {
	for {
		canal <- "pong" + time.Now().String()
		time.Sleep(time.Second*1) // Now, after (1 to print + 1 to pong) 1 seconds, a new pong was add into channel
		// ^ At the test, some prints register a different time interval - Searching for something
	}
}

func printPingPong(canal chan string) {
	for {
		mensagem := <-canal // <- Again, but now, data comes from channel canal to variable mensagem
		fmt.Println(mensagem)
		time.Sleep(time.Second*1) // Wait for 1 second
	}
}
