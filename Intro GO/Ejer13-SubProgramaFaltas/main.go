package main

import "fmt"

//Stalumnos estructura con los datos de los alumnos (nombre, apellido, estado[presen,ausente]).""
type Stalumnos struct {
	nombre, apellido string
	estado           bool
}

//Stdias estructura con arreglo de alumnos de tipo StAlumnos, ademas contiene los dias de la semana.
type Stdias struct {
	listAlumno [21]Stalumnos
	dias       string
}

func main() {

	var listadoAlumnos [21]Stalumnos
	var dias [5]Stdias
	var diaElec string
	var presencia string
	elecControl := "s"
	var elecIntDay int

	listadoAlumnos[0].nombre = "Ivan"
	listadoAlumnos[0].apellido = "Aguiar"

	listadoAlumnos[1].nombre = "Nehemias"
	listadoAlumnos[1].apellido = "Astrada"

	listadoAlumnos[2].nombre = "Gonzalo"
	listadoAlumnos[2].apellido = "Blotta"

	listadoAlumnos[3].nombre = "Vincent"
	listadoAlumnos[3].apellido = "Conace"

	listadoAlumnos[4].nombre = "Joa Pablo"
	listadoAlumnos[4].apellido = "Corre Parise"

	listadoAlumnos[5].nombre = "Miguel"
	listadoAlumnos[5].apellido = "De Robles"

	listadoAlumnos[6].nombre = "Mauro"
	listadoAlumnos[6].apellido = "Fuentes"

	listadoAlumnos[7].nombre = "Andrea"
	listadoAlumnos[7].apellido = "Juarez"

	listadoAlumnos[8].nombre = "Galileo"
	listadoAlumnos[8].apellido = "Luna"

	listadoAlumnos[9].nombre = "Celeste"
	listadoAlumnos[0].apellido = "Masmut"

	listadoAlumnos[9].nombre = "Cintia"
	listadoAlumnos[9].apellido = "Murashima"

	listadoAlumnos[10].nombre = "Malena"
	listadoAlumnos[10].apellido = "Murcia"

	listadoAlumnos[11].nombre = "Paola"
	listadoAlumnos[11].apellido = "Pesantez"

	listadoAlumnos[12].nombre = "Alan"
	listadoAlumnos[12].apellido = "Ra"

	listadoAlumnos[13].nombre = "Juan"
	listadoAlumnos[13].apellido = "Ramone"

	listadoAlumnos[14].nombre = "Alejandro"
	listadoAlumnos[14].apellido = "Santin"

	listadoAlumnos[15].nombre = "Lautaro"
	listadoAlumnos[15].apellido = "Toledo"

	listadoAlumnos[16].nombre = "Nahuel"
	listadoAlumnos[16].apellido = "Trucco"

	listadoAlumnos[17].nombre = "Ivo"
	listadoAlumnos[17].apellido = "Vucetic"

	listadoAlumnos[18].nombre = "Agustina"
	listadoAlumnos[18].apellido = "Zagame"

	listadoAlumnos[19].nombre = "Lucas"
	listadoAlumnos[19].apellido = "Abot"

	listadoAlumnos[20].nombre = "Nicolas"
	listadoAlumnos[20].apellido = "Tesone"

	// fmt.Println("Alumnos")
	// for i := 0; i < len(listadoAlumnos); i++ {
	// 	fmt.Println(listadoAlumnos[i].nombre)
	// 	fmt.Println(listadoAlumnos[i].apellido)
	// 	fmt.Println("--------------------")
	// }

	dias[0].dias = "Lunes"
	dias[0].listAlumno = listadoAlumnos

	dias[1].dias = "Martes"
	dias[1].listAlumno = listadoAlumnos

	dias[2].dias = "Miercoles"
	dias[2].listAlumno = listadoAlumnos

	dias[3].dias = "Jueves"
	dias[3].listAlumno = listadoAlumnos

	dias[4].dias = "Viernes"
	dias[4].listAlumno = listadoAlumnos

	// for i := 0; i < len(dias); i++ {
	// 	fmt.Println("Dia", dias[i].dias)
	// 	for x := 0; x < 21; x++ {
	// 		fmt.Println(dias[i].listAlumno[x])
	// 	}

	// }

	fmt.Println("Asistencias Skill Factory Golang.")

	fmt.Println("Dias de la semana:")
	for i := 0; i < len(dias); i++ {
		fmt.Println(dias[i].dias)
	}

ValidacionDia:
	{
		fmt.Println("Ingrese el dia de la semana que desea tomar asistencia.")
		fmt.Scanf("%s", &diaElec)

		if diaElec == "Lunes" {
			elecIntDay = 0
		} else if diaElec == "Martes" {
			elecIntDay = 1
		} else if diaElec == "Miercoles" {
			elecIntDay = 2
		} else if diaElec == "Jueves" {
			elecIntDay = 3
		} else if diaElec == "Viernes" {
			elecIntDay = 4
		}

		switch diaElec {
		case "Lunes":
			fmt.Println("Dia Lunes")
			fmt.Println()
			fmt.Println("Alumnos: ")
			for x := 0; x < len(dias[0].listAlumno); x++ {
				fmt.Println("---------------------")
				fmt.Println(dias[0].listAlumno[x].nombre)
				fmt.Println(dias[0].listAlumno[x].apellido)
				fmt.Println("---------------------")
				fmt.Println("'P'-Presente | 'A'-Ausente")
				fmt.Scan(&presencia)
				fmt.Println("---------------------")
				if presencia == "P" || presencia == "p" {
					dias[0].listAlumno[x].estado = true
				}
			}

		case "Martes":
			fmt.Println("Dia Martes")
			fmt.Println()
			fmt.Println("Alumnos: ")
			for x := 0; x < len(dias[1].listAlumno); x++ {
				fmt.Println("---------------------")
				fmt.Println(dias[1].listAlumno[x].nombre)
				fmt.Println(dias[1].listAlumno[x].apellido)
				fmt.Println("---------------------")
				fmt.Println("'P'-Presente | 'A'-Ausente")
				fmt.Scan(&presencia)
				fmt.Println("---------------------")
				if presencia == "P" || presencia == "p" {
					dias[1].listAlumno[x].estado = true
				}

			}
		case "Miercoles":
			fmt.Println("Dia Miercoles")
			fmt.Println()
			fmt.Println("Alumnos: ")
			for x := 0; x < len(dias[2].listAlumno); x++ {
				fmt.Println("---------------------")
				fmt.Println(dias[2].listAlumno[x].nombre)
				fmt.Println(dias[2].listAlumno[x].apellido)
				fmt.Println("---------------------")
				fmt.Println("'P'-Presente | 'A'-Ausente")
				fmt.Scan(&presencia)
				fmt.Println("---------------------")

				if presencia == "P" || presencia == "p" {
					dias[2].listAlumno[x].estado = true
				}
			}
		case "Jueves":
			fmt.Println("Dia Jueves")
			fmt.Println()
			fmt.Println("Alumnos: ")
			for x := 0; x < len(dias[3].listAlumno); x++ {
				fmt.Println("---------------------")
				fmt.Println(dias[3].listAlumno[x].nombre)
				fmt.Println(dias[3].listAlumno[x].apellido)
				fmt.Println("---------------------")
				fmt.Println("'P'-Presente | 'A'-Ausente")
				fmt.Scan(&presencia)
				fmt.Println("---------------------")

				if presencia == "P" || presencia == "p" {
					dias[3].listAlumno[x].estado = true
				}
			}
		case "Viernes":
			fmt.Println("Dia Viernes")
			fmt.Println()
			fmt.Println("Alumnos: ")
			for x := 0; x < len(dias[4].listAlumno); x++ {
				fmt.Println("---------------------")
				fmt.Println(dias[4].listAlumno[x].nombre)
				fmt.Println(dias[4].listAlumno[x].apellido)
				fmt.Println("---------------------")
				fmt.Println("'P'-Presente | 'A'-Ausente")
				fmt.Scan(&presencia)
				fmt.Println("---------------------")
				if presencia == "P" || presencia == "p" {
					dias[4].listAlumno[x].estado = true
				}
			}
		default:
			fmt.Println("Dia Incorrecto.")
			goto ValidacionDia

		}
	} //fin ciclo de validacion
	fmt.Println("Desea ver los presentes del dia? [s][n]")
	fmt.Scan(&elecControl)
	if elecControl == "s" {
		fmt.Println("---------------------")
		fmt.Println("\t", dias[elecIntDay].dias)
		fmt.Println("---------------------")
		for index := 0; index < 21; index++ {
			if dias[elecIntDay].listAlumno[index].estado == true {
				fmt.Println("---------------------")
				fmt.Println(dias[elecIntDay].listAlumno[index].nombre)
				fmt.Println(dias[elecIntDay].listAlumno[index].apellido)
				fmt.Println("---------------------")
			}

		}

	}

} //fin ciclo main
