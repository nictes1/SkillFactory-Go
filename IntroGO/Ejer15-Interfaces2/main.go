package main

import "fmt"

type serVivo interface {
	estoyVivo() bool
}

//Ianimal interfaz de tipo animal
type Ianimal interface { // se crean funciones en comunes para las diferentes especies en este caso
	respirar()
	comer()
	dormir()         ///funciones en comun
	especie() string ///es obligatorio usar todas las funciones de las interfaces.
}

type perro struct {
	cantPatas  int
	respirando bool //estructuras especificas a cada estructura
	comiendo   bool
	durmiendo  bool
	ladrando   bool
	vivo       bool
}

type gato struct {
	cantPatas  int
	respirando bool //estructuras especificas a cada estructura
	comiendo   bool
	durmiendo  bool
	maullando  bool
	vivo       bool
}

func (p *perro) respirar()       { p.respirando = true }
func (p *perro) comer()          { p.comiendo = true }
func (p *perro) dormir()         { p.durmiendo = true }
func (p *perro) especie() string { return "Perro" }
func (p *perro) estoVivo()       { p.vivo = true }

func (g *gato) respirar()       { g.respirando = true }
func (g *gato) comer()          { g.comiendo = true }
func (g *gato) dormir()         { g.durmiendo = true }
func (g *gato) especie() string { return "Gato" }
func (g *gato) estoyVivo()      { g.vivo = true }

//AnimalesRespirando printea los valores de los animales que esten respirando
func AnimalesRespirando(anim Ianimal) {
	anim.respirar()
	fmt.Printf("Soy un %s y ya estoy respirando \n", anim.especie())
}

func estoyVivo(v serVivo) bool {
	return v.estoyVivo()
}

func main() {
	gatito := new(gato)
	AnimalesRespirando(gatito)

	Perrito := new(perro)
	AnimalesRespirando(Perrito)
	Perrito.vivo = true

	fmt.Printf("Estoy vivo = %t", estoyVivo(Perrito))
}
