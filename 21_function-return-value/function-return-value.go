package main

import "fmt"

func sayHello(name string) string  {
	var result string
	if name == ""{
		result = "Hello Jhon Doe"
	}else{
		result = "Hello, "+name
	}
	
	return result
}


func main()  {
	result := sayHello("")
	fmt.Println(result)
	fmt.Println(sayHello("Rochi Eko"))

	name := sayHello("Pambudi")
	fmt.Println(name)
}