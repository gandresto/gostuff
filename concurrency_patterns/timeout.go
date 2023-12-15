package main

import (
	"fmt"
	"time"
)

func timeoutPerMessage() {
	const TIMEOUT_IN_MS = 650
	fmt.Printf("I get bored if I wait more than %d ms for a message\n", TIMEOUT_IN_MS)
	bob := boringGenerator("Bob!")
	for {
		select {
		case msg := <-bob:
			fmt.Println(msg)
		// time.After sends a message to a channel after a certain duration
		case <-time.After(time.Duration(TIMEOUT_IN_MS) * time.Millisecond):
			fmt.Println("You are too slow. Bye")
			return
		}
	}
}

func timeoutWholeConversation() {
	const TIMEOUT_IN_MS = 3000
	fmt.Printf("I get bored if the conversation last longer than %d ms\n", TIMEOUT_IN_MS)
	bob := boringGenerator("Bob!")
	timeout := time.After(time.Duration(TIMEOUT_IN_MS) * time.Millisecond)
	for {
		select {
		case msg := <-bob:
			fmt.Println(msg)
		// time.After sends a message to a channel after a certain duration
		case <-timeout:
			fmt.Println("This conversation lasted so much. Bye")
			return
		}
	}
}
