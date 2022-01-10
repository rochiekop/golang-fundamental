package main

import (
	"errors"
	"fmt"
)

func Devide(denominator int, numerator int) (int, error) {
	if numerator == 0 {
		return 0, errors.New("Devide by Zer0")
	} else {
		return (denominator / numerator), nil
	}
}

func main() {
	var err error = errors.New("Error")
	fmt.Println(err)

	result, err := Devide(100,34)
	if err == nil {
		fmt.Println("Result ", result)
	} else {
		fmt.Println("Error...", err.Error())
	}
}
