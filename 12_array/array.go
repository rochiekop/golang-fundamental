package main

import "fmt"


func main(){
	var names [4]string
	names[0] = "Rochi"
	names[1] = "Eko"
	names[3] = "Pambudi"
	names[2] = "Two"
	fmt.Println(string(names[0][0]))

	var values = [3]int{
		12,123,13,
	}
	fmt.Println(values[1])

	fmt.Println(len(values))
	values[2] = 100
	fmt.Println(values)

}