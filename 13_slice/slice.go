package main

import "fmt"

func main() {
	fmt.Println("Print")
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret", "April", "Mei", "Juni", "Juli",
		"Agustus", "September", "Oktober", "November", "Desember",
	}

	var slice1 = months[4:6]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// Change the array
	// months[10] = "Oktober New Value"
	months[5] = "Oktober New Value"
	fmt.Println(months)
	fmt.Println(slice1)

	days := [...]string{
		"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
	}
	fmt.Println(days)

	daySlice1 := days[5:]
	fmt.Println(daySlice1)
	daySlice1[0] = "New Saturday"
	daySlice1[1] = "New Sunday"
	fmt.Println(daySlice1)

	// Append
	daySlice2 := append(slice1, "New Day")
	daySlice2[0] = "Values"
	fmt.Println(daySlice2)
	fmt.Println(days)

	// newSlice

	newSlice := make([]string, 3, 4)
	newSlice[0] = "Slice1"
	newSlice[1] = "Slice2"
	newSlice[2] = "Slice3"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// copySlice
	fromSlice := days[:]
	var destinationSlice = make([]string, 3, 4)
	copy(destinationSlice, fromSlice)
	fmt.Println(destinationSlice)

	slice := []int{1,2,3,4,}
	array :=[...]int{1,2,3,4,}
	fmt.Println(slice)
	fmt.Println(array)
}
