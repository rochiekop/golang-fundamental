package main

import "fmt"

func recursiveFunction(number int) int {
	total := 1
	for i := number; i > 0; i-- {
		total *= i
	}
	return total
}

func recursiveFactorial(value int) int {
	if value == 1{
		return 1
	}else{
		fmt.Println(value)
		return value*recursiveFactorial(value-1)
	}
}

func main() {
	fmt.Println(recursiveFunction(10))
	fmt.Println(recursiveFactorial(10))
}
