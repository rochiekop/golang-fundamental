package main

import (
	"fmt"
	"os"
)


func main(){
	arg := os.Args
	fmt.Println("Arguments:")
	fmt.Println(arg)
	fmt.Println(arg[1])
	fmt.Println(arg[2])
	fmt.Println(arg[3])

	name, err := os.Hostname()

	if err == nil{
		fmt.Println("Hostname", name)
	}else{
		fmt.Println("Error", err)
	}
	// os.Setenv("APP-USERNAME", "root")
	// os.Setenv("APP-PASSWORD", "root")
	username := os.Getenv("APP-USERNAME")
	password := os.Getenv("APP-PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}