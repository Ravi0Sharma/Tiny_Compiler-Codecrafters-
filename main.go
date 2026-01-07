package main

import "fmt"

type token struct {
	kind  string
	value string
}

func tokenizer(input string) []token {
	input += "\n" // sentinel
	current := 0
	tokens := []token{}

	for current < len([]rune(input)) {
		char := string([]rune(input)[current])
		_ = char // temporary, until we use it
		break
	}

	return tokens
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
