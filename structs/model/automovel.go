package model

//Automovel é uma struct que armazena dados de um automovel
type Automovel struct {
	// PropertyName type `Tag[type of the tag:"Tag name"]`
	Wheels int `json:"number_of_wheels"` // Tags helps to transform struct representation
	Name string `json:"name"`
	value int // Not exported properties cant be tagged
	Potency float32 `json:"engine_potency"` 
}

// This methods belongs to the struct Automovel
// Inside () at func defition, is the scope of method or the struct -
// _ that his belongs
// Syntax: (variable name *Struct Name)
// Remeber: * at the begging represents the memory point to struct or a -
// referenced value

//SetValor define o valor do automovel
func (scope *Automovel) SetValue(value int) {
	scope.value = value
}

//GetValue retorna o valor do automovel
func (scope *Automovel) GetValue() int {
	return scope.value
}
