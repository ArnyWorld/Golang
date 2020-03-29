package main

import "fmt"

func main() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	/*For Clásico*/
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
	/*For infinito*/

	numero := 0
	for {
		fmt.Println("Continuo")
		fmt.Println("Ingrese el número secreto")
		fmt.Scanln(&numero)
		if numero == 100 {
			break
		}
	}

	/*For Continuo*/

	k := 0
	for k < 10 {
		fmt.Printf("\n valor de k: %d", k)
		if k == 5 {
			fmt.Printf("  Multiplicamos por 2 \n")
			k = k * 2
			fmt.Println("resultado ", k)
			continue
		}
		fmt.Printf("		Paso por aquí \n")
		k++
	}

	/*For controlado*/
	m := 0
RUTINA:
	for m < 10 {
		if m == 4 {
			m = m + 2
			fmt.Println("Voy a RUTINA y sumo 2")
			goto RUTINA
		}
		fmt.Printf("Valor de m: %d\n", m)
		m++
	}
}
