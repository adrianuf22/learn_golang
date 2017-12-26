package matematica

// Calc Executa um calculo passando os valores a e b
func Calc(fn func(int, int) int, a int, b int) int {
	return fn(a,b)
}

// @v6 Move to main to generico and exports
func Multiplicar(x int, y int) int {
	return x * y;
}

// @v7
//Dividir divide o valor a pelo b
// func Dividir(a int, b int) (resultado int) { // Using a named return value
// 	resultado = a / b // Named return value are automatic defined by Go, so just uses = to set the value
// 	return // Cause I using a named return value, an empty return must be used here
// }
// @v8
// Dividir divide o valor a pelo b
// Retorna o valor da divisão e o resto
func Dividir(a int, b int) (resultado int, resto int) { // To returns multiples values
	resultado = a / b
	resto = a % b
	return // Same as returns a single value
}
