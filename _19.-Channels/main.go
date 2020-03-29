package main

import (
	"fmt"
	"time"
)

func main() {
	/*Un channel es un espacio de memoria de dialogo en el que van a dialogar distintas rutinas
	y cuando se aloje un valor en ese espacio de memoria, la rutina que esta pidiendo un valor a acambio
	va actuar en consecuencia. Ese espacio de memoria debe ser de un tipo de dato.*/
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Llegue hasta aqui")

	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal1 chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 1000000000000; i++ {

	}
	final := time.Now()
	/*Sub es la funciòn que retorna la duraciòn*/
	canal1 <- final.Sub(inicio)
}
