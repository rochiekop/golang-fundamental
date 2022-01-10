package main

import "fmt"

func main()  {
	value := 0
	name := "Rochi"
	increment := func ()  {
		name := "Eko"
		value :=0
		fmt.Println("Increment")
		fmt.Println(name)	
		value++
	}
	
	increment()
	increment()
	print(name)
	fmt.Println(value)
}