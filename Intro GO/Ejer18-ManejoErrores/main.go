package main

//EjPanic ...
func EjPanic() {
	a := 1
	if a == 1 {
		panic("se encontro el valor 1")
	}
}

func main() {
	// archivo:="prueba.txt"
	// f, err := os.Open(archivo)

	// defer f.Close()

	// if err!= nil{
	// 	fmt.Println("error abriendo el archivo");
	// 	os.Exit(1)
	// }

	EjPanic()

}
