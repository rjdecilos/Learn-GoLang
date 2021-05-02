package main

type(
 Customer struct {
		FirstName string
		LastName string
		FullName string
	}
)

func getCustomers()(customers []Customer) {
	marcel := Customer{ FirstName: "Marcel", LastName: "Dempers"}

	customers = append(customers, marcel,
		Customer{ FirstName: "Bob", LastName: "Smith" },
		Customer{ FirstName: "Ben", LastName: "Spain" },
		Customer{ FirstName: "Aleem", LastName: "Jannmohamed" },
		Customer{ FirstName: "Jamie", LastName: "Le Notre" },
		Customer{ FirstName: "Victor", LastName: "Sankov" },
		Customer{ FirstName: "Adrian", LastName: "Oprea" },
	)

	return customers
}