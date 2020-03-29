package main

import "fmt"

func main() {
	numero := 5
	if numero == 5 {
		fmt.Println("Es 5")
	} else {
		fmt.Println("No es 5")
	}

	if numero = 6; numero == 5 {
		fmt.Println("Es 5")
	} else {
		fmt.Println("No es 5")
	}

	switch numero {
	case 1:
		fmt.Println("es 1")
	case 2:
		fmt.Println("es 2")
	case 3:
		fmt.Println("es 3")
	case 4:
		fmt.Println("es 4")
	case 5:
		fmt.Println("es 5")
	default:
		fmt.Println("Es otro número diferente")
	}

	switch numero = 4; numero {
	case 1:
		fmt.Println("es 1")
	case 2:
		fmt.Println("es 2")
	case 3:
		fmt.Println("es 3")
	case 4:
		fmt.Println("es 4")
	case 5:
		fmt.Println("es 5")
	default:
		fmt.Println("Es otro número diferente")
	}
}
