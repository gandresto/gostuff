package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boringWChan(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

// Generators are functions that return channels
func boringGenerator(msg string) <-chan string {
	c := make(chan string)
	// Launch the goroutine inside the function
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
	}()
	return c
}

func runBoringWithChan() {
	// channels are first class values, just like strings
	c := make(chan string)
	go boringWChan("boring!", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("Boring goroutine says: %q\n", <-c)
	}
	fmt.Println("Too boring. Main goroutine leaving.")
}

func runBoringGenerator() {
	c := boringGenerator("boring!")
	for i := 0; i < 5; i++ {
		fmt.Printf("Boring goroutine says: %q\n", <-c)
	}
	fmt.Println("Too boring. Main goroutine leaving.")
}

func bobAndAliceBoringConversation() {
	bob := boringGenerator("Bob message!")
	alice := boringGenerator("Alice message!")
	for i := 0; i < 5; i++ {
		// here they are both blocking each other
		// maybe bob is more talkative than alice, but still he has to wait for her
		fmt.Println(<-bob)
		// maybe alice is more talkative than bob, but still she has to wait for him
		fmt.Println(<-alice)
	}
	fmt.Println("You are both boring. Main goroutine leaving.")
}
