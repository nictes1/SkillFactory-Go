package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration) //time.duration tipo de dato.. puede devolver cualquier tipo de dato
	go bucle1(canal1)

	fmt.Println("Llegue hasta aqui")

	mensaje := <-canal1
	fmt.Println(mensaje)

}

func bucle1(miCanal chan time.Duration) bool {
	inicio := time.Now()
	for i := 0; i < 1000000; i++ {

	}
	final := time.Now()
	miCanal <- final.Sub(inicio) //duracion promedio, es lo que hay que devolver.
	//final sub inicio se le asiga al canal, ese valor es retornado
	return true
}
