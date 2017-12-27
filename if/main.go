package main

import (
	"strconv"
	"fmt"
)

var (
	meses int
	situacao bool
	pagamento string = "boleto"
	idUsuario int
	tratamento string
)

func main() {
	idUsuario = 1
	tratamento = "";

	meses = 0
	// @v1
	// if meses == 0 {
	// 	situacao = true
	// }

	// if meses > 0 {
	// 	fmt.Println("Meses em aberto: ", strconv.Itoa(meses))
	// }
	// Uses ; to delimiter variable definition and so check the condition
	// The variables defined here are only available to IF scope
	if descricao, status := haMesesEmAberto(meses); status {
		fmt.Println(descricao)
		situacao = status
	}
	// At compile, returns: undefined: descricao
	// fmt.Println(descricao);

	// ifElse
	if situacao {
		fmt.Println("A conta está bloqueada por falta de pagamento.")
		return
	} else {
		fmt.Println("A conta está ativa e pronta para uso.")
	}

	// The === is not used in GO, has no problems with types, so, just check the value
	// if pagamento === "boleto" { // unexpected =, expecting expression
	if pagamento == "boleto" {
		fmt.Println("Forma de pagamento escolhida: ", pagamento)
	}
	
	if naoEncontrado := usuarioExiste(idUsuario); naoEncontrado {
		fmt.Println("Usuário não encontrado!");
		return
	}

	codigo, nome, naoEncontrado := retornaUsuario(idUsuario)
	if naoEncontrado {
		usuarioNaoEncontrado();
		return
	}

	if registrarPagamento(codigo) {
		notificaCliente(nome, tratamento)
	}

	// Another way, now closing scope inside if
	if codigo, nome, naoEncontrado := retornaUsuario(idUsuario); !naoEncontrado {
		if registrarPagamento(codigo) {
			notificaCliente(nome, tratamento)
		}
	}
}

// @v2
func haMesesEmAberto(meses int) (descricao string, status bool) {
	if meses > 0 {
		descricao = "Meses em aberto: " + strconv.Itoa(meses)
		status = true
		return
	}

	descricao = "Contas em dia!"
	return
}

func usuarioExiste(id int) (naoEncontrado bool) {
	if id != 1 {
		naoEncontrado = false
		return
	}
	return
}

func retornaUsuario(id int) (codigo int, nome string, naoEncontrado bool) {
	if id == 1 {
		codigo = 1000
		nome = "Adriano"
		return
	}

	naoEncontrado = true
	return
}

func usuarioNaoEncontrado() {
	fmt.Println("Usuário não encontrado!");
}

func registrarPagamento(codigo int) bool {
	fmt.Println("Registrando pagamento...");
	fmt.Println("Pagamento efetuado pelo usuário #" + strconv.Itoa(codigo));
	fmt.Println("Registrado com sucesso!");
	
	return true
}

func retornaTratamento(tratamento string) (preposicao string, descricao string) {
	if tratamento == "sr" {
		preposicao = "do"
		descricao = "Sr."
		return
	}

	if tratamento == "sra" {
		preposicao = "da"
		descricao = "Sra."
		return
	}

	preposicao = "de"
	return
}

func notificaCliente(nome string, tratamento string) {
	preposicao, descricaoTratamento := retornaTratamento(tratamento)
	
	fmt.Printf("A conta %s %s %s foi recebida!\r\n", preposicao, descricaoTratamento, nome)
}
