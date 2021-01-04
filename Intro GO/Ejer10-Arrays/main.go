package main

import (
	"fmt"
)

var array [7]int
var matriz [2][3]int

func main() {
	fmt.Println(array)
	dim := cap(array)

	fmt.Println("Carga de una array.")
	for i := 0; i < dim; i++ {
		fmt.Println("Ingrese un numero")
		fmt.Scanf("%d ", &array[i])
	}

	for i := 0; i < dim; i++ {
		//fmt.Printf("[%d] \n", array[i])
		fmt.Println(array[i])
	}

	otherArray := [5]int{12, 3, 145, 76, 7}
	fmt.Println(otherArray)

	// fmt.Println("Carga Matriz")
	// filas := cap(matriz)
	// columnas := len(matriz)

	// fmt.Println(validos)
	// fmt.Println(columnas)
}
