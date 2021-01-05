package main

import "fmt"

func main() {
	potencia(1)
}

func potencia(num int) {

	if num > 1000 {
		return
	}
	fmt.Println(num)
	potencia(num * 2)
}
