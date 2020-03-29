package main

import "fmt"

/*Matriz con 5 FILAS y 3 COLUMNAS*/

func main() {
	const fila int = 5
	const columna int = 3
	var matriz [fila][columna]int

	matriz[0][0] = 1
	matriz[1][1] = 2
	matriz[2][2] = 3
	matriz[3][1] = 2
	matriz[4][0] = 1

	for i := 0; i < fila; i++ {
		for j := 0; j < columna; j++ {
			fmt.Printf(" %d", matriz[i][j])
		}
		fmt.Println()
	}
}
