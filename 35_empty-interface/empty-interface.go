package main

import "fmt"

func EmptyInterface(param int) interface{}{
	if param == 1 {
		return 1
	}else if param == 2 {
		return true
	}else {
		return "This is String"
	}
}

func main()  {
	var number interface{} = EmptyInterface(1)
	fmt.Println(number)
	fmt.Println(EmptyInterface(2))
	value := EmptyInterface(4)
	fmt.Println(value)
}