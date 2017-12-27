package model

//Automovel é uma struct que armazena dados de um automovel
type Automovel struct {
	// PropertyName type `Tag[type of the tag:"Tag name"]`
	Wheels int `json:"number_of_wheels"` // Tags helps to transform struct representation
	Name string `json:"name"`
	value int // Not exported properties cant be tagged
	Potency float32 `json:"engine_potency"` 
}

//SetValor define o valor do automovel
func (scope *Automovel) SetValue(value int) {
	scope.value = value
}

func (scope *Automovel) GetValue() int {
	return scope.value
}
