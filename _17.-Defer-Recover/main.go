package main

import "log"

func main() {
	/*Primero se ejecuta esta parte*/
	//archivo := "prueba.txt"
	//f, err := os.Open(archivo)
	/*Una declaración defer difiere la ejecución de una función hasta que la función circundante regrese.*/
	/*Al ùltimo ejecuta esta parte*/
	//defer f.Close()
	/*Segundo se ejecuta esta parte*/
	//if err != nil {
	//	fmt.Println("error abriendo el archivo")
	//	os.Exit(1)
	//}

	ejemploPanic()
}

func ejemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("ocurrio un error que genero Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("se encontro el valor de 1")
	}
}
