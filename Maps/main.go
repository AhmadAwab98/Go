package main

import "fmt"

type customer struct {
	name			string
	phoneNumber	string
	email   		string
	address			string
	age      		int
}

func main() {
	// Define a map to store employee information
	customerData := map[string] customer{
		"ahmad-khan": {
			name:     "Ahmad Khan",
			phoneNumber: "0336-3361234",
			email: "ahmad@gmail.com",
			address: "80-B Muslim Town Lahore",
			age:      30,
		},
		"fatima_malik": {
			name:     "Fatima Ahmad",
			phoneNumber: "0334-3344321",
			email: "fatima@gmail.com",
			address: "55-A Ichra Lahore",
			age:      25,
		},
		"ali_ahmad": {
			name:     "Ali Ahmad",
			phoneNumber: "0330-1258973",
			email: "ali@gmail.com",
			address: "220-G Johar Town Lahore",
			age:      45,
		},
	}

	// Print employee details
	for user_id, details := range customerData {
		fmt.Println("Customer: \n", user_id)
		fmt.Println("  Name: \n", details.name)
		fmt.Println("  Phone number: \n", details.phoneNumber)
		fmt.Println("  Email: \n", details.email)
		fmt.Println("  Address: \n", details.address)
		fmt.Println("  Age: \n", details.age)
		fmt.Println("-------------------------------")
	}
}