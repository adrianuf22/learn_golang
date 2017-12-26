package matematica

// Calc Executa um calculo passando os valores a e b
func Calc(fn func(int, int) int, a int, b int) int {
	return fn(a,b)
}

// @v6 Move to main to generico and exports
func Multiplicar(x int, y int) int {
	return x * y;
}
