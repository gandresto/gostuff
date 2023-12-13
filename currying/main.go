package main

import "fmt"

type formatterFunct func(string, string) string

// getLogger takes a function that formats two strings into
// a single string and return a function that formats two strings
// but prints the result instead of returning it
func getLogger(formatter formatterFunct) func(string, string) {
	return func(s1, s2 string) {
		fmt.Println(formatter(s1, s2))
	}
}

func commaDelimited(strA, strB string) string {
	return strA + "," + strB
}

func colonDellimited(strA, strB string) string {
	return strA + ":" + strB
}

func appendStr(strA, strB string) string {
	return strA + strB
}

func main() {
	getLogger(colonDellimited)("First", "Second")
	getLogger(commaDelimited)("First", "Second")
	getLogger(appendStr)("First", "Second")
}
