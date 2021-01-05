package main

import "fmt"

//Los middelwares son interceptores que permiten ejecutar intrucciones
//comunes a varias funciones que recibe y devuelven los mismo tipos de
//variables

func main() {

	fmt.Println("Inicio de programa")

	result := operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(5, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(4, 3)
	fmt.Println(result)

}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operacion")
		return f(a, b)
	}
}
