package main

import "fmt"

func sayHello(firstName string, lastName string){
	fmt.Println("Hello", firstName, lastName)
}

func main(){
	sayHello("Rochi", "Eko Pambudi")
	firstName := "Eko"
	lastName := "Khennedy"

	sayHello(firstName,lastName)
}