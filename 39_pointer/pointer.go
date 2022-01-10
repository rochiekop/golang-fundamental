package main

import "fmt"


type Address struct {
	city, province, country string

}

func ChangeAddressToIndonesia(address *Address)  {
	address.country = "Indonesia"
}

func main()  {
	address1 := Address{"Pemalang", "Central Java", "Indonesia"}
	fmt.Println(address1)

	address2 := address1
	address2.city = "Bandung"
	fmt.Println(address2)
	fmt.Println(address1)

	// Pointer (&)
	/* address2 := &address1
	address2.city = "Bandung"
	address2.province = "West Java"
	
	fmt.Println(address2)
	fmt.Println(address1)

	address2 := &address1
	address2.city = "Bandung"
	address2.province = "West Java"
	fmt.Println(address2)
	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	var address3 *Address = &address1

	*address3 = Address{"Jakarta", "DKI Jakarta", "Indonesia"}


	fmt.Println(address1) //Data not change
	fmt.Println(address2)
	fmt.Println(address3) 

	 */
	// Function New
	address3 := new(Address)
	address4 := address3

	address4.city = "Bandung"

	fmt.Println(address3)
	fmt.Println(address4)

	address5 := Address{"Pemalang","Central", ""}
	addressPointer := &address5
	ChangeAddressToIndonesia(addressPointer)
	fmt.Println(address5)

}