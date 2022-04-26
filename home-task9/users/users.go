package users

type Employee struct {
	CurrentAge int
}

type Customer struct {
	CurrentAge int
}

type User interface {
	Employee | Customer
	Age() int
}

func (e Employee) Age() int {
	return e.CurrentAge
}

func (c Customer) Age() int {
	return c.CurrentAge
}
