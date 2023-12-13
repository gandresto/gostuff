package main

import "fmt"

type person struct {
	name, lastName string
}

type fullTimeEmployee struct {
	salary float64
	person
}

type contractor struct {
	payPerHour  float64
	hoursWorked uint
	person
}

type employee interface {
	getName() string
	getSalary() float64
}

func (p person) getName() string {
	return p.name + " " + p.lastName
}

func (ft fullTimeEmployee) getSalary() float64 {
	return ft.salary
}

func (ct contractor) getSalary() float64 {
	return ct.payPerHour * float64(ct.hoursWorked)
}

func logSalary(e employee) {
	fmt.Printf("%s salary is %.2f\n", e.getName(), e.getSalary())
}

func main() {
	e1 := fullTimeEmployee{salary: 12384.1, person: person{name: "Andres", lastName: "Gonzalez"}}
	logSalary(e1)

	e2 := contractor{hoursWorked: 50, payPerHour: 123.0, person: person{name: "Itzel", lastName: "Mendoza"}}
	logSalary(e2)
}
