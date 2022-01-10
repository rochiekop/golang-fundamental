package main

import "fmt"


func main(){
	var lulusUjian = 90
	var lulusPresensi = 85

	fmt.Println(lulusPresensi >= 90 && lulusUjian >= 90)
	fmt.Println(lulusPresensi >= 80 && lulusUjian >= 90)
}