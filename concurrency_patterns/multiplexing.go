package main

import "fmt"

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	// First approach: launch one goroutine per channel

	// go func() {
	// 	for {
	// 		c <- <-input1
	// 	}
	// }()
	// go func() {
	// 	for {
	// 		c <- <-input2
	// 	}
	// }()

	// Second approach: one goroutine and select statement
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

func multiplexingExample() {
	bob := boringGenerator("Bob")
	ann := boringGenerator("Ann")
	c := fanIn(bob, ann)
	for i := 0; i < 14; i++ {
		// We could note that we could recieve more messages of one of them
		// but we eliminate the lockstep between them
		fmt.Println(<-c)
	}
	fmt.Println("You are both boring. Main goroutine leaving.")
}
