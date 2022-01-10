package main

import "fmt"


func main(){
	name := "Rochi"

	if name == "Rochi"{
		fmt.Println("Hello Rochi")
	}else if name == "Eko"{
		fmt.Println("Hello Eko")
	}else{
		fmt.Println("What your name ?")
	}

	if lengthName := len(name); lengthName >= 5{
		fmt.Println("Accepted")
	}else{
		fmt.Println("Name is to short")
	}
}