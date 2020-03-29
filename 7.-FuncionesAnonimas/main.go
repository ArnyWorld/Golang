package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("La suma de los números 5 y 7: %d\n", Calculo(5, 7))
	/*Redefiniendo Calculo-Ahora va restar*/
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Printf("La resta de los números 5 y 3: %d\n", Calculo(5, 3))
	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Printf("Dividimos los números 12 y 3: %d\n", Calculo(12, 3))

	Operaciones()
}

func Operaciones() {
	resultado := func() int {
		var a int = 10
		var b int = 7
		return a + b
	}
	fmt.Printf("La suma de los numeros es = %d", resultado())
}
