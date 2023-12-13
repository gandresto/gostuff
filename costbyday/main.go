package main

import "fmt"

type cost struct {
	day   int
	value float64
}

const totalDaysOfTheWeek = 7

func calculateCostsByDay(costs []cost) (costsByDay []float64) {
	costsByDay = make([]float64, totalDaysOfTheWeek)
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		costsByDay[cost.day] += cost.value
	}
	return costsByDay
}

func formatCostByDay(costsByDay []float64) (formatted string) {
	for _, cost := range costsByDay {
		formatted += fmt.Sprintf("%.2f ", cost)
	}
	return formatted
}

func main() {
	costs := []cost{
		{day: 0, value: 3.2},
		{day: 2, value: 1.2},
		{day: 3, value: 3.5},
		{day: 1, value: 3.2},
		{day: 2, value: 4.2},
		{day: 5, value: 9.2},
		{day: 2, value: 6.2},
		{day: 6, value: 3.7},
		{day: 0, value: 3.9},
		{day: 1, value: 1.2},
		{day: 2, value: 3.0},
		{day: 3, value: 1.0},
		{day: 4, value: 2.2},
		{day: 5, value: 4.9},
		{day: 6, value: 7.5},
	}

	fmt.Println(formatCostByDay(calculateCostsByDay(costs)))
}
