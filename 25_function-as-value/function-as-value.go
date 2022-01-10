package main

import "fmt"


func sayHello(param string) string  {
	return "Hello "+param
}


func main()  {
	var getSayHello = sayHello
	fmt.Println(getSayHello("Rochi Eko Pambudi"))
	setHello := getSayHello("Eko")
	fmt.Println(setHello)
}