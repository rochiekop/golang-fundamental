package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.Round(60.4))
	fmt.Println(math.Ceil(60.7))
	fmt.Println(math.Floor(60.7))
	fmt.Println(math.Max(600, 599.5))
	fmt.Println(math.Min(600, 599.5))

	fmt.Printf("%.2f", math.Cos(90/2))
	fmt.Println()
	fmt.Printf("%.1f", math.Log10(100))
	fmt.Println()
	fmt.Println(math.Exp(23))
}
