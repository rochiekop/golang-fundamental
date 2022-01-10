package main

import (
	"flag"
	"fmt"
)

func main() {
	var host *string = flag.String("host", "localhost", "Your your database host")
	var username *string = flag.String("username", "root", "Put your database username")
	var password *string = flag.String("password", "root", "Put your database password")
	var number *int = flag.Int("number", 101, "Put your number")

	flag.Parse()

	fmt.Println("Host", *host)
	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Number", number)
}
