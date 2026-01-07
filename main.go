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
	runes := []rune(input) // så vi inte bygger []rune(input) hela tiden

	for current < len(runes) {
		char := string(runes[current])

		// 1) Parentheses
		if char == "(" {
			tokens = append(tokens, token{kind: "paren", value: "("})
			current++
			continue
		}
		if char == ")" {
			tokens = append(tokens, token{kind: "paren", value: ")"})
			current++
			continue
		}

		// 2) Whitespace
		if char == " " || char == "\n" || char == "\t" || char == "\r" {
			current++
			continue
		}

		// 3) Numbers (commit 8)
		if isNumber(char) {
			value := ""
			for current < len(runes) && isNumber(string(runes[current])) {
				value += string(runes[current])
				current++
			}
			tokens = append(tokens, token{kind: "number", value: value})
			continue
		}

		// 4) Names (commit 9)
		if isLetter(char) {
			value := ""
			for current < len(runes) && isLetter(string(runes[current])) {
				value += string(runes[current])
				current++
			}
			tokens = append(tokens, token{kind: "name", value: value})
			continue
		}

		// 5) Unknown character error (commit 10)
		log.Fatalf("tokenizer: unexpected character %q at position %d", char, current)
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
