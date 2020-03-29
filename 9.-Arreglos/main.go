package main

import "fmt"

/* Forma de crear un vector por fuera de main
var tabla [10]int
*/

func main() {
	tabla := [10]int{5, 1, 2, 8, 7, 6, 4, 3, 0, 9}

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	/*
		Forma de asignar valores a un arreglo
		tabla[0]=5
		tabla[1]=1

	*/

}
