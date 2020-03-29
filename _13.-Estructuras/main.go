package main

import (
	"fmt"
	"time"

	/*importando el paquete creado*/
	us "./user"
)

/*Creando una estructura de acuerdo al paquete importado*/
type arni struct {
	us.Usuario
}

func main() {
	u := new(arni)
	u.AltaUsuario(1, "Arnaldo Muñoz", time.Now(), true)
	fmt.Println(u.Usuario)
}
