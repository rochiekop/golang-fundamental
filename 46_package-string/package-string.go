package main

import (
	"fmt"
	"strings"
)


func main(){
	fmt.Println(strings.Trim("           Rochi Eko Pambudi"," "))
	fmt.Println(strings.Contains("Rochi Eko Pambudi","Eko"))
	fmt.Println(strings.Contains("Rochi Eko Pambudi","eko"))
	fmt.Println(strings.Contains("Rochi Eko Pambudi","Dsd"))
	
	fmt.Println(strings.Split("Rochi Eko Pambudi"," "))

	fmt.Println(strings.ToLower("Rochi Eko Pambudi"))

	fmt.Println(strings.ToTitle("rochi eko pambudi"))
	fmt.Println(strings.ToUpper("Rochi Eko Pambudi"))
	
	fmt.Println(strings.ReplaceAll("Rochi Eko Pambudi", "Eko","Replace"))

}
