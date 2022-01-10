package main

import (
	"container/ring"
	"fmt"
	"strconv"
)


func main(){
	var data *ring.Ring = ring.New(5)

	// data.Value = "$Rochi"
	// fmt.Println(data)
	// data2 := data.Next()
	// data2.Value = "Eko"
	// fmt.Println(data2)

	for i:=0; i < data.Len(); i++{
		data.Value = "Data number "+ strconv.FormatInt(int64(i), 10)
		data = data.Next()
	}

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})
}