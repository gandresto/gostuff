package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Printf("%s %d\n", msg, i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func justRunBoring() {
	boring("boring!")
}

func runBoringWithGoroutine() {
	go boring("boring!")
}

func runBoringWithMainGoroutineWaiting() {
	go boring("boring!")
	fmt.Println("I'm listening")
	time.Sleep(time.Duration(2) * time.Second)
	fmt.Println("Too boring. Main goroutine leaving.")
}
