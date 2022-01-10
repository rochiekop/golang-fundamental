package main

import "fmt"

func getName()(firstName string, middleName string, lastName string)  {
	firstName = "Rochi"
	middleName = "Eko"
	lastName = "Pambudi"

	return
}

func main()  {
	a,b,c := getName()
	fmt.Println(a,b,c)
}