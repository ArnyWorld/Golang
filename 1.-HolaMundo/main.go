/*el valor de package debe ser identico al nombre del archivo ,en este caso el nombre
del archivo es main*/
package main

/*La libreria fmt nos permite mostrar algo por consola*/
import "fmt"

/*la función main es la función principal del programa.*/
func main() {
	fmt.Println("Hola Mundo")
}

/*Para ejecutar el programa main.go debe ejecutarse la instrucción

							go run main.go

Se debe tener el cuidado de estar ubicado dentro de la carpeta qeu contiene el
archivo main.go
Para generar el ejecutable (main.exe) se debe ejecutar la instrucción

							go build main.go

*/
