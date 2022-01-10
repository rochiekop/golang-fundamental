package main

import "fmt"

func main() {
	var name string = "Rochi Eko Pambudi"

	// type1
	switch name {
	case "Rochi":
		fmt.Println("Firstname", name)
	case "Eko Pambudi":
		fmt.Println("Lastname", name)
	default:
		fmt.Println("Name is correctly")
	}

	// type2
	switch lengthName := len(name); lengthName > 10 {
	case true:
		fmt.Println("Name is long")
	case false:
		fmt.Println("Name is short")
	}

	// type3
	lengthName := len(name)

	switch {
	case (lengthName > 5 && lengthName <= 10):
		fmt.Println("Name is short")
	case (lengthName > 10 && lengthName <= 20):
		fmt.Println("Name is normal")
	case (lengthName > 20 && lengthName <= 30):
		fmt.Println("Name is long")
	case lengthName > 30:
		fmt.Println("Name is too long")
	default:
		fmt.Println("Please input name")
	}

}
