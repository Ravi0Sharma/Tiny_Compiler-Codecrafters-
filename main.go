package main

import "fmt"

type token struct {
	kind  string
	value string
}

func tokenizer(input string) []token {
	return []token{}
}

func isNumber(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	return n >= '0' && n <= '9'
}

func isLetter(char string) bool {
	if char == "" {
		return false
	}
	n := []rune(char)[0]
	return n >= 'a' && n <= 'z'
}

func main() {
	input := "(add 2 (subtract 10 5))"
	fmt.Println(input)
}
