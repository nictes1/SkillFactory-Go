package main

import "log"

//EjPanic manejo de errores.
func EjPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrio un error que genero panic. %v \n", reco)
		}
	}()

	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}

func main() {
	EjPanic()
}
