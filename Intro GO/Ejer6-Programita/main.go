package main

import "fmt"

var num int

func main() {

	numSecreto := 9
	var primeraVez bool = true

	for {
		fmt.Println("Ingrese un numero")
		fmt.Scanln(&num)

		if num == numSecreto {
			fmt.Println("Adivinaste el numero")
			if primeraVez == true {
				primeraVez = false
				continue
			} else {
				break
			}

		} else {
			fmt.Println("No adivinaste el numero")
		}
	}

}
