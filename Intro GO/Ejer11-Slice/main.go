package main

import (
	"fmt"
)

var slice []int

func main() {
	// slice = []int{2, 5, 8, 11}
	// fmt.Println(slice)
	// slice = append(slice, 14)
	// fmt.Println(slice)

	// fmt.Println("Otro SLICE")
	// otherSlice()

	otherSlice2()
	otherSlice3()

}

func otherSlice() {
	elementos := [5]int{1, 2, 3, 4, 5}
	porcion := elementos[2:5]
	fmt.Println(elementos)
	fmt.Println(porcion)
}

func otherSlice2() {
	element := make([]int, 6, 20)
	fmt.Println(element)

	fmt.Printf("Large: %d. Capacity: %d", len(element), cap(element))

}

func otherSlice3() {
	elements := make([]int, 0, 0)
	for i := 0; i < 10; i++ {
		elements = append(elements, i)
	}
	fmt.Printf("Large: %d. Capacity: %d", len(elements), cap(elements))
}

//prueba cap y make
// func main() {
// 	slice := []int{1, 2, 3, 4}
// 	copia := make([]int, len(slice))

// 	copy(copia, slice)

// 	fmt.Println(slice)
// 	fmt.Println(copia)

// }
