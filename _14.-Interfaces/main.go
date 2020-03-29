package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
	estaVivo() bool
}
type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}
type vegetal interface {
	ClasificacionVegetal() string
}

/*Genero Humano*/
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
	vivo       bool
}
type mujer struct {
	/*De esta forma se hereda las propiedades del hombre*/
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) sexo() string {
	if h.esHombre == true {
		return "Hombre"
	} else {
		return "Mujer"
	}
}
func (h *hombre) estaVivo() bool {
	return h.vivo
}
func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("Soy un/a %s, y ya estoy respirando \n", hu.sexo())
}

/******Genero Animal****/
type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) EsCarnivoro() bool {
	return p.carnivoro
}
func (p *perro) estaVivo() bool {
	return p.vivo
}

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("Soy un animal y estoy respirando")
}
func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() == true {
		return 1
	}
	return 0
}

/****Ser Vivo******/
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}
func main() {
	/*
		Arni := new(hombre)
		Arni.esHombre = true
		HumanosRespirando(Arni)

		Lendy := new(mujer)
		HumanosRespirando(Lendy)

	*/

	totalCarnivoros := 0
	Dogo := new(perro)
	Dogo.carnivoro = true
	Dogo.vivo = true
	AnimalesRespirar(Dogo)
	totalCarnivoros = +AnimalesCarnivoros(Dogo)

	fmt.Printf("Total Carnivoros %d \n", totalCarnivoros)

	fmt.Printf("Estoy vivo %t", estoyVivo(Dogo))

}
