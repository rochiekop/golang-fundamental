package main

import "fmt"


func sayHello()(string,string,string)  {
	return "Rochi", "Eko", "Pambudi"
}

func main()  {
	firstName,_,lastName := sayHello()

	fmt.Println(firstName)
	fmt.Println(lastName)
}