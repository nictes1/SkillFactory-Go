package main

import (
	"fmt"
)

var calculo func(int, int) int = func(numero1 int, numero2 int) int {
	suma := numero1 + numero2
	return suma
}

func main() {
	fmt.Printf("Sumo 5+5 = %d \n", calculo(5, 5))

	calculo = func(numero1 int, numero2 int) int {
		resta := numero1 - numero2
		return resta
	}
	fmt.Printf("resto 10+5 = %d \n", calculo(10, 5))
}
