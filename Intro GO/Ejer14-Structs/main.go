package main

// import (
// 	"fmt"
// 	"time"
// )

// type usuario struct {
// 	ID        int
// 	Nombre    string
// 	FechaALta time.Time
// 	status    bool
// }

// func main() {
// 	User := new(usuario)
// 	User.ID = 10
// 	User.Nombre = "Nicolas"
// 	User.FechaALta = time.Now()
// 	fmt.Println(User)
// }

import (
	"fmt"
	"time"

	us "./User" //directorio donde se encuentra guardado la libreria.
)

type alumno struct{
	
}


//array de structs 

var arrayStructs[30] [nombre de la estructura] (perro)

for i:=0;i>len(arrayStructs);i++{
	perri := new(perro) //nombre de la estructura
	arrayStructs[i]= *perri
}

fmt.Println(arrayStructs)