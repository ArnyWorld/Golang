package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("Ingrese número 1:")
	/*La función Scanln nos ayuda a leer un dato por consola, el parametro que recibe es una variable
	donde se almacenara el dato leido, esta variable debe estar acompañado de un puntero (&)*/
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese número 2:")
	fmt.Scanln(&numero2)

	fmt.Println("Descripción :")

	/*bufio es una librería que nos ayudara a leer texto por consola y se ayudara de la función NewScanner
	que como parametro utiliza la librería os para utilizar la función Stdin que permite leer una cadena (string)
	y debe ser un dato de entrada (input)
	*/
	scanner := bufio.NewScanner(os.Stdin)
	/*Evalua si se esta ingresando algun dato por consola*/
	if scanner.Scan() {
		/*Si se ingresa un adto lo almacena en la variable leyenda*/
		leyenda = scanner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
