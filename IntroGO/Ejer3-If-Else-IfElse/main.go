package main

import "fmt"

var estado bool

func main() {
	estado = true

	if estado == true {
		fmt.Println("Es verdadero")
	} else {
		fmt.Println("Es falso")
	}

}
