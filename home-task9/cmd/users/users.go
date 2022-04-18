package main

import (
	"home-task9/users"
)

func main() {
	var emp1 = users.Employee{CurrentAge: 21}
	var emp2 = users.Employee{CurrentAge: 22}
	var emp3 = users.Employee{CurrentAge: 23}

	var cus1 = users.Customer{CurrentAge: 24}
	var cus2 = users.Customer{CurrentAge: 19}
	var cus3 = users.Customer{CurrentAge: 25}

	maxAge(emp1, emp2, emp3)
	maxAge(cus1, cus2, cus3)
}

func maxAge[T users.User](peoples ...T) int {
	var max = 0

	if len(peoples) == 0 {
		return max
	}

	for _, human := range peoples {
		if human.Age() > max {
			max = human.Age()
		}
	}

	return max
}
