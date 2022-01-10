package main

import (
	"golang-fundamental/41_helper"
	"fmt"
)

func main()  {
	helper.SayHello("Eko Pambudi")
	// helper.sayHello("Name") //Error
	fmt.Println(helper.Application)
	// fmt.Println(helper.private) //Error
}