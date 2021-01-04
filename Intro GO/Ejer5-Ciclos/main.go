package main

import "fmt"

func main() {
	i := 1
	for i < 7 {
		fmt.Println(i)
		i++
	}

	//OR
	for i := 1; i <= 7; i++ {
		fmt.Println(i)
	}

}
