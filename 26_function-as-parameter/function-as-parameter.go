package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilte(name string, filter Filter) {
	filterName := filter(name)
	fmt.Println("Hello ", filterName)
}

func Filtered(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func countNumber(name string) string {
	number := len(name)

	if number >= 10 {
		return name
	} else {
		return "Jhon Doe"
	}
}

func main() {
	sayHelloWithFilte("Anjing", Filtered)
	filter := Filtered
	sayHelloWithFilte("Rochi Eko Pambudi", filter)

	sayHelloWithFilte("Rochi Eko", countNumber)
}
