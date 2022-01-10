package main

import "fmt"

type User struct {
	name, address string
	age           int
	isMale        bool
}


func (user User) sayHello(name string){
	fmt.Println("Hello ", user.name, "My Name is", name)
}

func main(){
	rochi := User{name:"Rochi"}
	rochi.sayHello("Eko")
}