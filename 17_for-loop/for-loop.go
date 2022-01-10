package main

import "fmt"

func main(){

	for index := 1; index <= 10; index++{
		fmt.Println("Data ke ", index)
	}

	//Slice

	// Type 1
	days := []string{
		"Sunday", "Monday", "Tuesday",
	}

	for i:=0; i < len(days); i++{
		fmt.Println("Index ", i ,"=", days[i])
	}

	// Type 2

	for i, day := range days{
		fmt.Println("Index", i , "=",day)
	} 

	for _, day := range days{
		fmt.Println(day)
	} 

	// Map
	var movies map[string]string = make(map[string]string)
	movies["title"] = "Avenger"
	movies["year"] = "2019"
	movies["rating"] = "8.9"
	movies["desc"] = "Desc"
	
	for key, value :=range movies{
		fmt.Println(
			key,":",value,
		)
	}
}