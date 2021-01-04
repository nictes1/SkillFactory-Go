package main

import (
	"fmt"
	"strings"
)

func main() {
	var NombreCompleto string
	var Acronimo string
	// fmt.Println("Ingrese su nombre completo: ")
	// fmt.Scanf("%s", &NombreCompleto)
	NombreCompleto = "Organizacion de las Naciones Unidas"

	fmt.Println(len(NombreCompleto))

	palabras := strings.Fields(NombreCompleto)

	for _, palabra := range palabras {

		pLetra := string(palabra[0])
		if pLetra == strings.ToUpper(pLetra) {
			Acronimo += strings.ToUpper(pLetra)
		}

	}
	fmt.Println(Acronimo)
}
