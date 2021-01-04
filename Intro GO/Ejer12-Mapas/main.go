package main

import "fmt"

func main() {

	countrys := make(map[string]string, 4)

	countrys["Brasil"] = "Brasilia"
	countrys["Argentina"] = "Buenos Aires"
	countrys["Paraguay"] = "Asuncion"
	countrys["Colombia"] = "Bogota"
	fmt.Println(countrys)

	Tournament := map[string]int{

		"River Plate": 45,
		"Lanus":       23,
		"Sao Pablo":   30,
	}
	fmt.Println(Tournament)

	delete(Tournament, "Lanus") ///para borrar un equipo, en este caso lanus.

	fmt.Println(Tournament)

	for team, points := range Tournament {
		fmt.Printf("El equipo %s, tiene %d puntos \n", team, points)
	}

	points, existe := Tournament["River Plate"]

	if existe == true {
		fmt.Printf("Existe el equipo, y tiene %d puntos \n", points)
	} else {
		fmt.Println("No existe el equipo")
	}

}
