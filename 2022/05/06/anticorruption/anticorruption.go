package anticorruption

import (
	"github.com/MarioCarrion/videos/2022/05/06/fahrenheit"
	"github.com/MarioCarrion/videos/2022/05/06/kelvin"
)

func Kelvin(country string) float64 {
	// C = K − 273.15
	return kelvin.Calculate(country) - 273.15
}

func Fahrenheit(country string) float64 {
	// C = (5.0/9.0) * (F-32)
	return (5.0 / 9.0) * (fahrenheit.New(country).Value() - 32)
}
