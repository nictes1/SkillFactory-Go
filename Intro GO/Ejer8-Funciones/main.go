package main

//Librerias
import "fmt"

//Variables Globales

func main() {
	//funcion con un solo retorno
	numero := 5
	resultado := primeraFuncion(numero)
	fmt.Println(resultado)

	//funciones con mas de un retorno
	num, estado := segundaFuncion(numero)
	fmt.Println(num)
	fmt.Println(estado)

	//funcion con parametros indefinidos (...)
	sumTotal := calculo(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println(sumTotal)
}

func primeraFuncion(numero int) int {
	return numero * 2
} //funcion con una sola devolucion

func primeraFuncionOther(numero int) (doble int) { //Llamamos doble al return.
	doble = numero * 2 //asginamos el resultado al return llamado doble
	return             //devuelve sin parametros ya que el return se llama doble y ya esta asignado
} //funcion con una sola devolucion

//funcion con dos returns.
func segundaFuncion(numero int) (int, bool) {
	if numero == 5 {
		return 5, true
	}
	return 10, false
}

//fn con parametros indefinidos
func calculo(numero ...int) int {
	total := 0
	for _, valor := range numero { //range devuelve dos cosas, la subpos y el valor
		total = total + valor
	}
	return total
}
