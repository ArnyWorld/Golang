package main

import "fmt"

func main() {
	/*Los vectores dinamicos son conocidos como slices*/
	matriz := []int{2, 3, 5}
	fmt.Println(matriz)

	variante2()
	variante3()
	variante4()
}
func variante2() {
	elementos := [5]int{1, 2, 3, 4, 5}
	/*Los dos puntos se refieren a una porción, en este
	caso se quiere separar los elementos desde la posición 3
	para adelante y desde la posicion 0 hasta la 4*/
	porcion := elementos[3:]
	porcion2 := elementos[:4]

	fmt.Println(porcion)
	fmt.Println(porcion2)

}
func variante3() {
	/*Creando un slice de tipo entero, longitud, capacidad(reserva de memoria)*/
	elementos := make([]int, 5, 20)
	/*len nos ayuda a conocer la longitud o tamaño de un arreglo, cap nos ayuda a conocer la capacidad máxima de un arreglo*/
	fmt.Printf("Largo %d, Capacidad %d", len(elementos), cap(elementos))
}
func variante4() {
	/*La reserva de cantidad automatica se basa en la base binaria
	si se almacena un valor en el vector, se reservara la capacidad de 2
	si se almacena 8 se reservara 16
	si se almacena 100 se reservara 128. Siempre se reservara ppotencias de 2*/
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\n Largo %d, Capacidad %d", len(nums), cap(nums))
}
