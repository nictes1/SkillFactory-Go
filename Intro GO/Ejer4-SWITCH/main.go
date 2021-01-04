package main

import "fmt"

var numero int

func main() {
	numero = 4

	switch numero {
	case 1:
		fmt.Println("Vale 1")
	case 2:
		fmt.Println("Vale 2")
	default:
		fmt.Println("Vale mas que 2")
	}

	///Or

	//switch numero = 1; numero {
	//case 1:
	//	fmt.Println("Vale 1")
	//case 2:
	//	fmt.Println("Vale 2")
	//default:
	//	fmt.Println("Vale mas que 2")
	//}
}
