package main

import (
	"log"
	"os"
)

//EjPanic manejo de errores.
func EjPanic() {

	//abre el archivo webserver.log para escritura
	f, err := os.OpenFile("webserver.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Fatal(err)
	}

	//cierra cuando termine la funcione del main
	defer f.Close()

	//asocia el manejador del archivo al log
	log.SetOutput(f)

	//agregan 10 lineas al archivo log
	for i := 0; i <= 10; i++ {
		log.Printf("Error en la linea %v", i)
	}

	/*
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
	*/

}

func main() {
	EjPanic()
}
