package model

type Imovel struct {
	Size int
	Garage bool
	Name string
	value int
}

func (i *Imovel) SetValue(value int) {
	i.value = value
}
