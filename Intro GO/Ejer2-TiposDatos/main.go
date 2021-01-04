package main

import (
	"fmt"
	"strconv"
)

var numero, numeroaux, numerototal int
var texto string
var estado bool

func main() {

	var numeroMain int

	var moneda int
	//conversion de float a char
	///numeroEj := int(moneda) //forma de casteo
	moneda = 30
	TxtCovertido := strconv.Itoa(moneda) //Int to Ascii
	fmt.Println(TxtCovertido)

	untexto := "Hola Nicolas, Como andas?" //:= Declada la variable y el compilador detecta que tipo de dato es

	numero = 4
	numeroaux = 2

	fmt.Println(numero)
	fmt.Println(numeroaux)

	numerototal = numero + numeroaux

	fmt.Println(numerototal)
	fmt.Println(texto)
	fmt.Println(estado)
	fmt.Println(numeroMain)
	fmt.Println(untexto)

	estado = true

	mostrarEstado()

}

func mostrarEstado() {
	fmt.Println(estado)
}
