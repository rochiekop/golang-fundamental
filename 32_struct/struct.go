package main

import "fmt"

type User struct {
	name, address string
	age           int
	isMale        bool
}

func main() {
	// create struct 1
	var user1 User
	user1.name = "Rochi Eko Pambudi"
	user1.address = "Bandung, West Java"
	user1.age = 24
	user1.isMale = true

	fmt.Println(user1)
	fmt.Println(user1.name)

	// create struct 2

	user2 := User{
		name: "Rochi Eko Pambudi",
		address: "Bandung, West Java",
		age: 24,
		isMale: true,
	}
	fmt.Println(user2)
	fmt.Println(user2.address)

	// create struct 3
	// Consider the order of parameter

	user3 := User{"Rochi Eko Pambudi", "Bandung, West Java", 24, true}
	fmt.Println(user3)

}
