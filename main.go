package main

import "fmt"
import "strings"

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	// Trim whitespace, lowercase, then split on any whitespace
	text = strings.TrimSpace(strings.ToLower(text))
	if text == "" {
		return []string{}
	}
	return strings.Fields(text)
}
