package main

import "fmt"

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
}

type gato struct {
	cantPatas  int
	respirando bool //estructuras especificas a cada estructura
	comiendo   bool
	durmiendo  bool
	maullando  bool
}

func (p *perro) respirar()       { p.respirando = true }
func (p *perro) comer()          { p.comiendo = true }
func (p *perro) dormir()         { p.durmiendo = true }
func (p *perro) especie() string { return "Perro" }

func (g *gato) respirar()       { g.respirando = true }
func (g *gato) comer()          { g.comiendo = true }
func (g *gato) dormir()         { g.durmiendo = true }
func (g *gato) especie() string { return "Gato" }

//AnimalesRespirando printea los valores de los animales que esten respirando
func AnimalesRespirando(anim Ianimal) {
	anim.respirar()
	fmt.Printf("Soy un %s y ya estoy respirando \n", anim.especie())
}

func main() {
	gatito := new(gato)
	AnimalesRespirando(gatito)

	perrito := new(perro)
	AnimalesRespirando(perrito)
}
