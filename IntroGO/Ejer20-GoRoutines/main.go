package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	go miNombreLento("Nicolas")
	go miNombreLento("Tesone")
	time.Sleep(10000 * time.Millisecond)
	fmt.Println("Estoy aqui")
	var x string
	fmt.Scanln(x)
}

func miNombreLento(nombre string) {
	letras := strings.Split(nombre, "")
	for _, letras := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letras)
	}
}
