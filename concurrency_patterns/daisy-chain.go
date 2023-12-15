package main

import (
	"fmt"
)

func incrementAndPassToLeft(left chan<- int, right <-chan int) {
	left <- 1 + <-right
}

func daisyChain() {
	const N_CHANNELS = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < N_CHANNELS; i++ {
		right = make(chan int)
		go incrementAndPassToLeft(left, right)
		left = right
	}

	go func(c chan<- int) {
		c <- 1
	}(right)
	fmt.Print(<-leftmost)
}
