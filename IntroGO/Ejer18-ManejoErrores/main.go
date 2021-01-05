package main

import (
	"fmt"
	"os" //paquete para abrir el archivo
)

//EjPanic ...
func EjPanic() {
	a := 1
	if a == 1 {
		panic("Se encontro el valor 1")
	} else {
		fmt.Println("El programa no entro al PANIC")
	}
}

func main() {
	archivo := "prueba.txt" //no existe el archivo, va arrojar error
	f, err := os.Open(archivo)

	defer f.Close() //defer se ejecuta antes de finalizar cualquier programa.

	if err != nil {
		fmt.Println("error abriendo el archivo")
		os.Exit(1) //el numero puede ir cambiando en el caso de necesitar varios exits.
	}

	///////

	EjPanic()

}
