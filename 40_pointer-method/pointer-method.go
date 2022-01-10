package main

import "fmt"

type Man struct {
	Name string
}

func (male *Man) Married() {
	male.Name = "Mr." + male.Name
}

func main() {
	diki := Man{"Diki"}
	diki.Married()
	fmt.Println(diki.Name)
}
