package main

import (
	"container/list"
	"fmt"
)

func main() {
	list := list.New()

	list.PushBack("Rochi")
	list.PushBack("Eko")
	list.PushBack("Pambudi")
	fmt.Println(list.Front().Value)
	fmt.Println(list.Back().Value)
	fmt.Println(list.Front().Prev())
	fmt.Println(list.Back().Next())

	// Iterate Sequence
	for value := list.Front(); value != nil; value = value.Next() {
		fmt.Println(value.Value)
	}

	// Reverse
	for value := list.Back(); value != nil; value = value.Prev(){
		fmt.Println(value.Value)
	}
}
