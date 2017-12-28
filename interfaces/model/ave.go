package model

//Ave é um modelo para animais do tipo ave
type Ave struct { // Has no "implements" or any keyword to chain the struct with an interface
	Name string
	flying bool
	comida string
}

// In GO, the interface is identified by methods defined in the struct
// Struct must provide the required interface methods and so it will be implemented

//Cacareja retorna a frase famosa da ave galinha
func (scope Ave) Cacareja() string { // Don't use * or a pointer
	//     ^ When implements an interface, scope wasn't a reference or a pointer
	// See below: Using a pointer
	return "Cocoricó!!!"
}

//Grasna retorna a frase famose da ave pato
func (scope Ave) Grasna() string {
	return "Quack Quack!!!"
}

// But a point can be used too
func (scope *Ave) TipoComida(comida string) {
	scope.comida = comida;
}

//Voa faz a ave voar ou sair do chão
func (scope *Ave) Voa() {
	scope.flying = true
}
