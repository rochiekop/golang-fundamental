package main

import "fmt"

func newMap(name string, address string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	var person map[string]string = newMap("", "Bandung")
	if person == nil {
		fmt.Println("Data Empty..")
	} else {
		fmt.Println(person)
	}
}
