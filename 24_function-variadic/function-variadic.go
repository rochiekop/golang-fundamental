package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	var number int = sumAll(10, 11, 8, 4, 1, 20)
	fmt.Println(number)

	slice := []int{10, 20, 5, 15, 25}
	number = sumAll(slice...)
	fmt.Println(number)
}
