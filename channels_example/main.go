package main

import (
	"fmt"
	"math/rand"
	"time"
)

var logger = getLoggerWithDate()

func getLoggerWithDate() func(string, ...any) (int, error) {
	return func(format string, a ...any) (n int, err error) {
		timestamp := time.Now().Format("2006-01-02T15:04:05.999Z")
		prefix := fmt.Sprintf("[%s]", timestamp)
		return fmt.Printf(prefix+" "+format, a...)
	}
}

func makeRequest(id string, requestChan chan<- string) {
	const MAX_RANDOM_VALUE = 123
	randomInt := rand.Intn(1500)
	time.Sleep(time.Duration(randomInt) * time.Millisecond)
	logger("Process%s its ready to send data to chan. \n", id)
	requestChan <- id
	logger("Process%s sent some data to the chan. \n", id)
}

func processRequest(ch <-chan string) {
	logger("\t\t\t\t\t\tReading data from chan: " + <-ch + "\n")
}

func main() {
	const ARR_SIZE = 5
	processIds := [ARR_SIZE]string{"1", "2", "3", "4", "5"}

	requestChan := make(chan string)
	for _, processId := range processIds {
		logger("Creating process%s. \n", processId)
		go makeRequest(processId, requestChan)
	}

	for i := 0; i < ARR_SIZE; i++ {
		processRequest(requestChan)
	}
}
