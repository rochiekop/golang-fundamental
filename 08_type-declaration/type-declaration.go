package main

import "fmt"


func main(){
	type noHp string
	type noKtp string

	var no_ktp noKtp = "0493458579545"
	var no_hp noHp = "0985654"

	fmt.Println(no_ktp)
	fmt.Println(no_hp)
	fmt.Println(noHp("0954357"))
}