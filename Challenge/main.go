package main

import "fmt"

type Alumnos struct {
	dni              int32
	nombre, apellido string
	estado           bool
}

type Materia struct {
	alum    []Alumnos
	materia string
}

var Curso []Materia
var ArrayAlumnos []Alumnos

func main() {

	var Alum1, Alum2, Alum3, Alum4, Alum5 Alumnos
	Alum1 = NewStudent("Franco", "Armani", 42492172)
	Alum2 = NewStudent("Marcelo", "Gallardo", 36947182)
	Alum3 = NewStudent("Enzo", "Perez", 37564324)
	Alum4 = NewStudent("Nacho", "Fernandez", 34886020)
	Alum5 = NewStudent("Matias", "Suarez", 37884912)

	ArrayAlumnos = append(AddStudent(Alum1)) 
	ArrayAlumnos = append(AddStudent(Alum2)) 
	ArrayAlumnos = append(AddStudent(Alum3)) 
	ArrayAlumnos = append(AddStudent(Alum4)) 
	ArrayAlumnos = append(AddStudent(Alum5)) 
	
	MostrarStudents()
	DeleteStudent(42492172)
	MostrarStudents()

}

//se le pasan los datos del alumno por parametro y retorna un dato del tipo alumno.
func NewStudent(nombre string, apellido string, dni int32) Alumnos {
	return Alumnos{
		nombre:   nombre,
		apellido: apellido,
		dni:      dni,
		estado:   true,
	}

}

//se necesita la estructura de los alumnos con los datos previamente cargados.
func AddStudent(alumnito Alumnos) {
	existe := ExisteDni(alumnito.dni) ///validacion de que no existe el alumno

	if existe == true {
		fmt.Println("Ya existe ese alumno.")

	} else {
		arrayAlumnos = append(arrayAlumnos, Alumnos{ //si no existe el alumno se lo agrega a la el arreglo
			nombre:   alumnito.nombre,
			apellido: alumnito.apellido,
			dni:      alumnito.dni,
			estado:   alumnito.estado,
		})
	}
}

//funcion de validacion para la carga de los alumnos al slice
func ExisteDni(dni int32) bool {
	i := 0
	for i = 0; i < len(arrayAlumnos); i++ {
		if dni == arrayAlumnos[i].dni {
			return true
		}
	}
	return false
}

//no borra un alumno sino que le modifica el estado. Baja Logica-
func DeleteStudent(dni int32) {
	i := 0
	pos := 0
	flag := 0

	for i = 0; i < len(arrayAlumnos); i++ {
		if arrayAlumnos[i].dni == dni {
			flag = 1
			pos = i
		}
	}
	if flag == 1 {
		arrayAlumnos[pos].estado = false
	}
}

//Muestra los alumnos que esten activos(true). si estan con la baja logica, los alumnos no se mostraran.
func MostrarStudents() {
	i := 0
	validos := len(arrayAlumnos)

	for i = 0; i < validos; i++ {
		if arrayAlumnos[i].estado {
			fmt.Println(arrayAlumnos[i])
		}
	}
}

//fn que agrega la matria al curso. Revisar.
func AddMaterias(nombre string) {
	Existe := ValidCurso(nombre)
	if Existe == true {
		Curso = append(Curso, Materia{
			materia: nombre,
			alum:    nil,
		})
	} else {
		fmt.Println("El curso que desea ingresar ya existe.")
	}

}

//validar si existe o no el curso.
func ValidCurso(nombre string) bool {
	flag := false
	for i := 0; i < len(Curso); i++ {
		if Curso[i].materia == nombre {
			flag = true
		}
	}
	return flag
}

func AddAlumnos2Curso(alumnito Alumnos, nombreCurso string) {
	Existe := ValidCurso(nombreCurso)

	if Existe == true {
		for i := 0; i < len(Curso); i++ {
			if Curso[i].materia == nombreCurso {
				Curso = append(Curso[i], Materia{
					materia: nombreCurso
					al : append(alum)
				})
			}
		}

	}

}
