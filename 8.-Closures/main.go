package main

import "fmt"

func main() {
	tablaDel := 2
	MiTabla := Tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(MiTabla())
	}
}
func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia = secuencia + 1
		return numero * secuencia
	}
}
