package main

import "fmt"

func main() {

	planets := make(map[string]float64)
	var Planeta string
	var seconds float32
	var Flag bool

	planets["Earth"] = 31557600
	planets["Mercury"] = 0.2408467
	planets["Venus"] = 0.61519726
	planets["Mars"] = 1.8808158
	planets["Jupiter"] = 11.862615
	planets["Saturn"] = 29.447498
	planets["Uranus"] = 84.016846
	planets["Neptune"] = 164.79132

Validacion:
	{
		fmt.Println("Ingrese el planeta con el cual desea calcular su edad.")
		fmt.Println(planets)
		fmt.Scanln(&Planeta)

		for plan := range planets {
			if plan == Planeta {
				Flag = true
			}
		}

		if Flag != true {
			goto Validacion
		}

		fmt.Println("Ingrese los segundos")
		fmt.Scanln(&seconds)

		switch Planeta {
		case "Earth":
			return seconds / 31557600
		case "Mercury":
			return Years(seconds, 0.2408467)
		case "Venus":
			return Years(seconds, 0.61519726)
		case "Mars":
			return Years(seconds, 1.8808158)
		case "Jupiter":
			return Years(seconds, 11.862615)
		case "Saturn":
			return Years(seconds, 29.447498)
		case "Uranus":
			return Years(seconds, 84.016846)
		case "Neptune":
			return Years(seconds, 164.79132)
		}
	}

}

func Years(seconds float64, Orbital float64) float64 {
	return seconds / (3155760 * Orbital)
}
