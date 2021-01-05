package main

import "fmt"

//Alumnos estructura que contiene los campos en comun de un alumno
type Alumnos struct {
	dni              int32
	nombre, apellido string
	estado           bool
}

//Materia Contiene los alumnos y el nombre de la materia correspondiente
type Materia struct {
	alum    []Alumnos
	materia string
}

//Curso array  que contiene todos los campos Alumnos y materias
var Curso []Materia

//ArrayAlumnos se almacenan todos los alumnos
var ArrayAlumnos []Alumnos

func main() {

	Alum1 := NewStudent("Franco", "Armani", 42492172)
	Alum2 := NewStudent("Marcelo", "Gallardo", 36947182)
	Alum3 := NewStudent("Enzo", "Perez", 37564324)
	Alum4 := NewStudent("Nacho", "Fernandez", 34886020)
	Alum5 := NewStudent("Matias", "Suarez", 37884912)

	AddStudent(Alum1)
	AddStudent(Alum2)
	AddStudent(Alum3)
	AddStudent(Alum4)
	AddStudent(Alum5)

	AddMaterias("Matematica")
	AddMaterias("Literatura")

	AddAlumnos2Curso(Alum1, "Matematica")
	AddAlumnos2Curso(Alum2, "Matematica")
	AddAlumnos2Curso(Alum3, "Matematica")
	AddAlumnos2Curso(Alum4, "Matematica")
	AddAlumnos2Curso(Alum5, "Literatura")
	fmt.Println(Curso)

}

//NewStudent se le pasan los datos del alumno por parametro y retorna un dato del tipo alumno.
func NewStudent(nombre string, apellido string, dni int32) Alumnos {
	return Alumnos{
		nombre:   nombre,
		apellido: apellido,
		dni:      dni,
		estado:   true,
	}

}

//AddStudent se necesita la estructura de los alumnos con los datos previamente cargados.
func AddStudent(alumnito Alumnos) {
	existe := ExisteDni(alumnito.dni) ///validacion de que no existe el alumno

	if existe == true {
		fmt.Println("Ya existe ese alumno.")

	} else {
		ArrayAlumnos = append(ArrayAlumnos, Alumnos{ //si no existe el alumno se lo agrega a la el arreglo
			nombre:   alumnito.nombre,
			apellido: alumnito.apellido,
			dni:      alumnito.dni,
			estado:   alumnito.estado,
		})
	}
}

//ExisteDni funcion de validacion para la carga de los alumnos al slice
func ExisteDni(dni int32) bool {
	i := 0
	for i = 0; i < len(ArrayAlumnos); i++ {
		if dni == ArrayAlumnos[i].dni {
			return true
		}
	}
	return false
}

//DeleteStudent no borra un alumno sino que le modifica el estado. Baja Logica-
func DeleteStudent(dni int32) {
	i := 0
	pos := 0
	flag := 0

	for i = 0; i < len(ArrayAlumnos); i++ {
		if ArrayAlumnos[i].dni == dni {
			flag = 1
			pos = i
		}
	}
	if flag == 1 {
		ArrayAlumnos[pos].estado = false
	}
}

//MostrarStudents Muestra los alumnos que esten activos(true). si estan con la baja logica, los alumnos no se mostraran.
func MostrarStudents() {
	i := 0
	validos := len(ArrayAlumnos)

	for i = 0; i < validos; i++ {
		if ArrayAlumnos[i].estado {
			fmt.Println(ArrayAlumnos[i])
		}
	}
}

//AddMaterias fn que agrega la matria al curso. Revisar.
func AddMaterias(nombre string) {
	Existe := ValidCurso(nombre)
	if Existe == true {
		fmt.Println("El curso que desea ingresar ya existe.")
	} else {
		Curso = append(Curso, Materia{
			materia: nombre,
		})
	}

}

//ValidCurso validar si existe o no el curso.
func ValidCurso(nombre string) bool {
	flag := false
	for i := 0; i < len(Curso); i++ {
		if Curso[i].materia == nombre {
			flag = true
		}
	}
	return flag
}

//AddAlumnos2Curso : agrega a los alumnos al curso que se desee
func AddAlumnos2Curso(alumnito Alumnos, nombreCurso string) {
	Existe := ValidCurso(nombreCurso)
	ExisteEnCurso := false

	if Existe == true {
		ExisteALum := ExisteDni(alumnito.dni)

		if ExisteALum == true {
			for i, _ := range Curso {
				if Curso[i].materia == nombreCurso {
					for z, _ := range Curso[i].alum {
						if Curso[i].alum[z].dni == alumnito.dni {
							fmt.Println("El alumno ya se encuentra cargado")
							ExisteEnCurso = true
						}
					}
					if ExisteEnCurso != true {
						Curso[i].alum = append(Curso[i].alum, Alumnos{
							dni:      alumnito.dni,
							nombre:   alumnito.nombre,
							apellido: alumnito.apellido,
						})
					}
				}
			}
		}
	}
}
