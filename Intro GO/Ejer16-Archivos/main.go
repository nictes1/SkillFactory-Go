package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leoArchivo("./archivo.txt")
	// leoArchivo2()
	escribirArchivo("./archivoNuevo.txt")
	leoArchivo("./archivoNuevo.txt")
	escribirArchivo2("./archivoNuevo.txt")
}

func leoArchivo(dirArchi string) {
	archivo, err := ioutil.ReadFile(dirArchi)
	if err != nil { //evalua que alla encontrado el archivo.
		fmt.Println("Hubo un error.")
	} else {
		fmt.Println(string(archivo)) //funcion de conversion de los string
	}
}

func leoArchivo2(dirArchi string) {
	archivo, err := os.Open(dirArchi) // se abre el archivo
	if err != nil {                   //evalua que alla encontrado el archivo.
		fmt.Println("Hubo un error.")
	} else {
		scanner := bufio.NewScanner(archivo) //se crea un scanner con la libreria "bufio"
		for scanner.Scan() {                 //escanea linea por linea, no importa la cantidad
			registro := scanner.Text()
			fmt.Println(registro + "\n")
		}
	}
	archivo.Close() // siempre hay que finalizar el archivo.
}

func escribirArchivo(dirArchi string) {
	archivo, err := os.Create(dirArchi) // crea un archivo nuevo, y si hay un archivo con el mismo nombre lo pisa.
	if err != nil {                     //evalua que alla encontrado el archivo.
		fmt.Println("Hubo un error.")
	} else {
		fmt.Fprintln(archivo, "Funcionoo") // fprintln escribe en archivo
	}
	archivo.Close()
}

func escribirArchivo2(dirArchi string) {
	if AbrirAgregar(dirArchi, "\nHola buenas tardes, funciono bien.") == false {
		fmt.Println("Hubo un error. Al agregar datos")
	}

}

//AbrirAgregar abre el archivo y guarda un texto en el archivo. retorna true para exitoso.
func AbrirAgregar(dirArchi string, texto string) bool {
	arch, err := os.OpenFile(dirArchi, os.O_RDWR|os.O_APPEND, 0664) // os.O_WrONLY se abre para escritura  || os.O_Rw ONLY lectura unicamente ||  os.O_RdWr lecutra/escritura // 0664 permisos de lecto/escritura.
	if err != nil {
		fmt.Println("Hubo un error. Al abrirse el archivo")
		return false
	}
	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Hubo un error. En la escritura del archivo")
		return false
	}
	arch.Close()
	return true
}
