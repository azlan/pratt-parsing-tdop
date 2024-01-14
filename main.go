package main

import (
	"fmt"
)

var precedences = map[string]int{
	// EQ:       EQUALS,
	// NOT_EQ:   EQUALS,
	// LT:       LESSGREATER,
	// GT:       LESSGREATER,
	"PLUS":     SUM,
	"MINUS":    SUM,
	"SLASH":    PRODUCT,
	"ASTERISK": PRODUCT,
	"LPAREN":   CALL,
}

func main() {
	input := "(2 + 3 * (3 +1))"
	getTokens(input)
	fmt.Printf("%+v\n", tokens)
	sum := parseExpression(LOWEST)
	fmt.Println("final", sum)
}
