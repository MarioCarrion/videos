package main

import (
	"fmt"

	"gitlab.com/MarioCarrion/examples/2025/bonus/v1"
)

func main() {
	fmt.Println("Mario", bonus.Default.Percentage("Mario"))
	fmt.Println("Ruby", bonus.Default.Percentage("Ruby"))

	//-

	bonus.Default = bonus.NewCalculator(map[string]int{"Mario": 200, "Ruby": 300})
	bonus.DefaultValue = 500

	fmt.Println("Mario", bonus.Default.Percentage("Mario"))
	fmt.Println("Ruby", bonus.Default.Percentage("Ruby"))
}
