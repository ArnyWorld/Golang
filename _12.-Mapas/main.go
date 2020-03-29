package main

import "fmt"

func main() {
	paises := make(map[string]string, 5)
	fmt.Println(paises)
	paises["Bolivia"] = "Sucre"
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)

	/*Otra forma de crear maps*/
	campeonato := map[string]int{
		"The Strongest": 39,
		"Bolivar":       35,
		"Always Ready":  34,
		"San Jose":      30,
	}
	fmt.Println(campeonato)
	campeonato["Nacional Potosi"] = 25
	fmt.Println(campeonato)
	campeonato["Nacional Potosi"] = 55
	fmt.Println(campeonato)
	delete(campeonato, "Nacional Potosi")
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Mineiro"]
	fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
}
