package main

import "fmt"

func sum(nums ...float64) float64 {
	total := 0.0
	for i := 0; i < len(nums); i++ {
		total += nums[i]
	}
	return total
}

func main() {
	slice := []float64{1, 2, 3, 4, 5, 6}
	fmt.Printf("%.2f", sum(slice...))
}
