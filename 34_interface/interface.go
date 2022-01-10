package main

import "fmt"

type HasName interface {
	getName() string
}

func sayHello(hasName HasName) {
	fmt.Println("Hello ", hasName.getName())
}

type Person struct {
	name string
}

type Animal struct {
	name string
	isMamal bool
}

func (person Person) getName() string{
	return person.name
}

func (animal Animal) getName() (string) {
	return animal.name
}


func main() {

	person := Person{name: "Rochi"}
	sayHello(person)

	animal := Animal{name: "Cow", isMamal: true}
	sayHello(animal)
}
