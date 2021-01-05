package main

import (
	"fmt"
	"strings"
)

func main() {
	palabra := "animal"
	var inversa string

	for i := len(palabra) - 1; i >= 0; i-- {
		ultLetra := string(palabra[i])
		inversa += strings.ToUpper(ultLetra)
	}

	fmt.Println(inversa)
}
