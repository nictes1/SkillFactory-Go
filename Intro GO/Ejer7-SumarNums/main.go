package main

import "fmt"

func main() {

	var numero1, numero2, suma int

	fmt.Println("Ingrese un numero")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese otro numero")
	fmt.Scanln(&numero2)

	suma = numero1 + numero2

	fmt.Print("La suma de los dos valores es: ", suma)
}

/// OTHER EXAMPLE

// import (
// 	"fmt"
// 	"bufio"
//  "os"
// )
// func main() {

// 	var numero1, numero2 int
// 	var leyenda string

// 	fmt.Println("Ingrese un numero")
// 	fmt.Scanln(&numero1)

// 	fmt.Println("Ingrese otro numero")
// 	fmt.Scanln(&numero2)

// 	escaner := bufio.NewScanner(os.Stdin)
// 	if
// }
