package main

import (
	"fmt"

	"gitlab.com/MarioCarrion/examples/2025/bonus/v2"
)

func main() {
	fmt.Println("Mario", bonus.Default().Percentage("Mario"))
	fmt.Println("Ruby", bonus.Default().Percentage("Ruby"))

	// Compiler error:
	//-
	// bonus.Default = bonus.NewCalculator(map[string]int{"Mario": 200, "Ruby": 300}) // because is function
	// bonus.DefaultValue = 500  // because is constant
}
