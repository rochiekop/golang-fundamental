package main

import "fmt"

type Blocked func(string) bool

func registerUser(name string, blackList Blocked){
	if blackList(name){
		fmt.Println("Your Are Blocked as", name)
	}else{
		fmt.Println("Welcome", name)
	}
}

// func blackListAdmin(name string) bool {
// 	return name == "Admin"
// }

// func blackListRoot(name string) bool{
// 	return name == "Root"
// }

func main(){
	blackList := func(name string) bool {
		return name == "Root"
	}
	// registerUser("Root", blackListRoot)
	// registerUser("Root", blackListAdmin)
	registerUser("Root", blackList)

	registerUser("Admin", func(s string) bool {
		return s == "Root"
	})

	registerUser("User", func(s string) bool {
		return s == "Admin"
	})

	registerUser("Root", func(s string) bool {
		return s == "User"
	})
}