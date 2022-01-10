package main

import "fmt"

func main(){
	var name1 = "Eko"
	var name2 = "Pambudi"

	var rest1 bool = name1 == name2
	var rest2 bool = name1 != name2
	fmt.Println(rest1)
	fmt.Println(rest2)

	var num1 = 100
	var num2 = 230

	fmt.Println(num1 > num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 < num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
}