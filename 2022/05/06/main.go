package main

import (
	"fmt"

	"github.com/MarioCarrion/videos/2022/05/06/anticorruption"
	"github.com/MarioCarrion/videos/2022/05/06/fahrenheit"
	"github.com/MarioCarrion/videos/2022/05/06/kelvin"
)

func main() {
	fmt.Println("fahrenheit", fahrenheit.New("COUNTRY1").Value())
	fmt.Println("kelvin", kelvin.Calculate("COUNTRY2"))

	fmt.Println("anticorruption", anticorruption.Fahrenheit("COUNTRY1"))
	fmt.Println("anticorruption", anticorruption.Kelvin("COUNTRY2"))
}
