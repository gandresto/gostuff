package main

import "fmt"

func passByValue(m string) {
	fmt.Println("Pass by value")
	fmt.Println("Value: ", m)
	fmt.Println("Address: ", &m)
}

func passByReference(m *string) {
	fmt.Println("Pass by Reference")
	fmt.Println("Value: ", *m)
	fmt.Println("Address: ", m)
}

func main() {
	// the * syntax defines a pointer
	// var p *int

	// The & operator generates a pointer to its operand
	myStr := "Hello"
	pToMyStr := &myStr
	fmt.Println(pToMyStr)

	passByValue(myStr)
	passByReference(&myStr)
}
