package main

import "fmt"

var result int

func main() {
	/*Un middleware son interceptores, estos permiten
	ejecutar instrucciones comunes o varias funciones
	que reciben y devuelven los mismos tipos de variables*/

	fmt.Println("Inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(6, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(20, 5)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
