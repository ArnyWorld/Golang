package main

import "fmt"

func main() {
	/*fmt.Println(uno(15))
	numero, estado := dos(5)
	fmt.Println(numero)
	fmt.Println(estado)*/
	fmt.Println(Calculo(5, 46))
	fmt.Println(Calculo(2, 23, 45, 68))
	fmt.Println(Calculo(5))
	fmt.Println(Calculo(5, 46, 20, 45, 78))

}

/*Se devolvera un número entero*/
func uno(numero int) (z int) {
	z = numero * 2
	return
}

/*Devolvera mas de un valor*/
func dos(numero int) (int, bool) {
	if numero == 1 {
		return 5, false
	} else {
		return 10, true
	}
}

/*Funciones  con N parámetros*/
func Calculo(numero ...int) int {
	total := 0
	for item, num := range numero {
		total = total + num
		fmt.Printf("\n Item %d \n", item)
	}
	return total
}
