package main

import (
	/*La libreria strconv nos permitirá realizar una conversión de número a texto*/
	"fmt"
	"strconv"
)

/*Si se quiere crear variables globales, estas variables deben ir fuera de la función principal*/
var numero int
var texto string
var status bool

func main() {
	/*En GO existen tres tipos de variable
	la variables numéricas, string y booleanas
	De esta manera se crean variables con tipo de dato automatico, al asignar un dato a la variable
	estas reconoceran que datos se le estan asignando y automaticamente tendrán un tipo de dato
	*/
	numero1, numero2, numero3, texto := 2, 3, 5, "Soy una cadena"
	var moneda int64 = 0
	/*Conversion de número a texto*/
	texto = strconv.Itoa(int(moneda))
	fmt.Println(numero1)
	fmt.Println(numero2)
	fmt.Println(numero3)
	fmt.Println(texto)
}
