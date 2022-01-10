package main

import (
	"fmt"
	"strconv"
)

func main(){
	boolean, err := strconv.ParseBool("false")
	if err == nil {
		fmt.Println(boolean)
	} else{
		fmt.Println(err.Error())
		
	}

	Int, err := strconv.ParseInt("1000", 10, 16)

	if err == nil {
		fmt.Println(Int)
	} else {
		fmt.Println(err.Error())
	}

	stringNumber:= strconv.FormatInt(10000, 10)
	fmt.Println(stringNumber)

	
	stringNumberConv, _ := strconv.Atoi("1001")
	fmt.Println(stringNumberConv)

	intString:= strconv.Itoa(20000)
	fmt.Println(intString)

}