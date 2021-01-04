package main

import "fmt"

func main() {
	var año int
	fmt.Println("Ingrese un año para calular si es bisiesto o no")
	fmt.Scanln(&año)

	flag := deterAñoBisiesto(año)

	if flag == true {
		fmt.Printf("EL año %d, es un año biciento", año)
	} else {
		fmt.Printf("EL año %d, No es un año biciento", año)
	}

}

func deterAñoBisiesto(año int) bool {
	var flag bool

	if año%4 == 0 && año%100 != 0 || año%400 == 0 {
		flag = true
	}

	return flag
}
