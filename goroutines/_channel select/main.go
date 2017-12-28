package main

import (
	"fmt"
	"time"
)

var sms chan string
var irc = make(chan string) // Shorthand like a charm

func main() {
	sms = make(chan string) // Without initializated, nothing will happen at terminal
	// irc = make(chan string) // Or just use the shorthand

	go ping(sms)
	go pong(sms)
	go ying(irc)
	go yiang(irc)
	go printChannel()

	var input string
	fmt.Scanln(&input)
}

func ping(canal chan string) {
	for {
		canal <- "ping"
	}
}

func pong(canal chan string) {
	for {
		canal <- "pong"
	}
}

func ying(canal chan string) {
	for {
		canal <- "ying"
	}
}

func yiang(canal chan string) {
	for {
		canal <- "yiang"
	}
}

func printChannel() {
	var mensagem string
	for {
		select { // Like a switchCase
		case mensagem = <-sms: // Match the type and set the value to mensagem
			fmt.Println("From SMS: ", mensagem)
		case mensagem = <-irc:
			fmt.Println("From IRC: ", mensagem)
		}
		// @v2
		// time.Sleep(time.Second*1)
		time.Sleep(time.Second) // For just one second, *1 is not necessary cause the value of Second is 1s, so, when multiples to Second*2 (Behind the scenes: 1s*2) results 2s
	}
}
