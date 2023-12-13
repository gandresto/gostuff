package main

import "fmt"

func adder() func(int) int {
	acc := 0
	return func(a int) int {
		acc += a
		return acc
	}
}

type emailBill struct {
	costInPennies int
}

func test(bill []emailBill) {
	countAgg, costAgg := adder(), adder()
	for _, b := range bill {
		fmt.Printf("You have sent %d messages and have cost you $%d \n", countAgg(1), costAgg(b.costInPennies))
	}
}

func main() {
	test([]emailBill{
		{10},
		{42},
		{56},
		{113},
		{95},
		{8},
		{35},
	})
}
