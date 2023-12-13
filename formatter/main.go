package main

import "fmt"

func applyFormatter(msg string, formatter func(string) string) string {
	return formatter(msg)
}

func addSignature(msg string) string {
	return msg + " Regards. Andrés González"
}

func addGreeting(msg string) string {
	return "Hello! " + msg
}

func main() {
	fmt.Println(applyFormatter("Un mensaje shido.", addGreeting))
	fmt.Println(applyFormatter("Otro mensaje shido.", addSignature))
}
