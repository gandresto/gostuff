package main

import (
	"fmt"
)

type menu struct {
	description string
	f           func()
}

func main() {
	options := map[string]menu{
		"1":  {"Boring random sleep time program", justRunBoring},
		"2":  {"Boring with goroutine stops when main gorotine ends", runBoringWithGoroutine},
		"3":  {"Boring with goroutine stops after 2 secs", runBoringWithMainGoroutineWaiting},
		"4":  {"Boring with channel stops when revieves 5 messages", runBoringWithChan},
		"5":  {"Boring generator stops when revieves 5 messages", runBoringGenerator},
		"6":  {"Boring conversation stops when revieves 5 messages from each person", bobAndAliceBoringConversation},
		"7":  {"Multiplexing example", multiplexingExample},
		"8":  {"Timeout per message", timeoutPerMessage},
		"9":  {"Timeout whole conversation", timeoutWholeConversation},
		"10": {"Daisy chain of 100,000 goroutines", daisyChain},
		"0":  {"Just exit the program 🏃", func() {}},
	}

	pFunction, err := executeUserMenu(options)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	(*pFunction)()

}

func executeUserMenu(options map[string]menu) (*func(), error) {
	fmt.Println("Select an option to be executed.")
	for k, v := range options {
		fmt.Printf("%-3s%s\n", k, v.description)
	}

	var selectedOption string
	for {
		_, err := fmt.Scan(&selectedOption)

		if err != nil {
			return nil, nil
		}

		option, ok := options[selectedOption]

		if ok {
			return &option.f, nil
		}
		fmt.Println("Invalid option. Please select again.")
	}
}
