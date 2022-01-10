package main

import (
	"fmt"
)

func random() interface{} {
	return false
}

func main() {
	result := random()
	// fmt.Println(result)
	// resultString := result.(string)
	// fmt.Println(resultString)

	// resultInt := result.(int)
	// fmt.Println(resultInt) //Panic

	switch value := result.(type) {
	case string:
		fmt.Println("Type Value", value, "is String")
	case int:
		fmt.Println("Type Value", value, "is int")
	default:
		fmt.Println("Type Value", value, "is Boolean")
	}
}
