package main

import "fmt"

func endApp() {
	errorValue := recover()
	if errorValue != nil {
		fmt.Println("Whoops ", errorValue)
	}
	fmt.Println("End Application")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("Application ERROR..")
	}
	// Wrong Code
	// errorValue := recover()
	// fmt.Println("Whoops ", errorValue)
	fmt.Println("Application Running ...")
}

func main() {
	// runApp(false)
	runApp(true)
	fmt.Println("Rochi Eko Pambudi")
}
