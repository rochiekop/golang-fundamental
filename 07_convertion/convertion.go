package main

import "fmt"


func main(){
	var nilai32 int32 = 128
	var nilai64 int64 = int64(nilai32)
	var nilai18 int8 = int8(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai18)


	var name = "Eko Pambudi"
	var e = name[0]
	var eString = string(e)
	fmt.Println(eString)
}