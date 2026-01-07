package main

import "fmt"

type token struct {
	kind  string
	value string
}

func main() {
	input := "(add 2 (subtract 10 5))"
	fmt.Println(input)
}
