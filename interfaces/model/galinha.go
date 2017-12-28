package model

//Galinha representa uma ave do tipo galinha
type Galinha interface {
	Cacareja() string
	// @v2 Interfaces & pointer
	TipoComida(comida string)
}
