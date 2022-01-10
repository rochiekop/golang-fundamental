package main

import (
	// _ "golang-fundamental/43_database"
	"golang-fundamental/43_database"
	"fmt"
)

func main()  {
	conn := database.GetDatabase()
	fmt.Println(conn)
}