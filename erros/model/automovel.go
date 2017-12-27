package model

import (
	"errors" // To create error
)

type Automovel struct {
	Name string
	Wheels int
	Potency float32
	value int
}

/*
ErrValorDoAutomovelMuitoBaixo é um erro que é apresentado
ao definir um valor muito baixo a um automovel
*/
var ErrValorDoAutomovelMuitoBaixo = errors.New("O valor informado é muito baixo para um automóvel") // Creates a new Error object/struct, I dont know the right name, yet

/*
ErrValorDoAutomovelMuitoAlto é um erro que é apresentado
ao definir um valor muito alto a um automovel
*/
var ErrValorDoAutomovelMuitoAlto = errors.New("O valor informado é muito alto, é um carro ou um foguete?")

//SetValue define o valor do automovel
func (scope *Automovel) SetValue(value int) (err error) { // Assign the return
	err = nil
	if value < 15000 {
		err = ErrValorDoAutomovelMuitoBaixo // Uses the new Error
		return
	}

	if value > 5000000 {
		err = ErrValorDoAutomovelMuitoAlto
		return
	}

	scope.value = value
	return // Now, return is necessary cause, a return value was assigned
}
