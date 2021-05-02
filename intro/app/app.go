package main

import (
	"fmt"
)

func main() {
	customers := getCustomers()

	for _, customer := range customers {
		fmt.Println(customer)
	}
}

func getData() (customers []string) {
	
	customers = []string{ "Marcel Dempers", "Bob Smith", "John Smith" }
	
	customers = append(customers, "Ben Spain")
	customers = append(customers, "Aleem Jannmohamed")
	customers = append(customers, "Jamie Le Notre")
	customers = append(customers, "Victor Sankov")
	customers = append(customers, "Adrian Oprea")

	for _, customer := range customers {
		
		fmt.Println(customer)
	}
	return customers
}