package main

import "fmt"

func main(){
	person := map[string]string{
		"name" : "Rochi",
		"address" : "Bandung",
	}

	fmt.Println(person)
	
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	person["gender"] = "Male"
	fmt.Println(person)
	fmt.Println(len(person))

	var movies map[string]string = make(map[string]string)
	movies["title"] = "Avenger"
	movies["year"] = "2019"
	movies["rating"] = "8.9"
	movies["desc"] = "Desc"
	fmt.Println(movies)
	delete(movies,"desc")
	fmt.Println(movies)

}