package main

type Planeta string

const AñosTierra float64 = 31557600

var años map[Planeta]float64 = map[Planeta]float64{

	"Tierra":   1.0,
	"Mercurio": 0.2408467,
	"Venus":    0.61519726,
	"Marte":    1.8808158,
	"Jupiter":  11.862615,
	"Saturno":  29.447498,
	"Urano":    84.016846,
	"Neptuno":  164.79132,
}

func Calculo(segs float64, plan Planeta) float64 {
	return segs / AñosTierra / años[plan]
}

func main() {
	//..
}
