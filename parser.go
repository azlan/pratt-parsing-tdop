package main

import (
	"fmt"
	"log"
	"strconv"
)

func parseExpression(precedence int) int64 {
	tok := tokens[idx]
	var value int64
	if tok.kind == "NUMBER" {
		value, _ = strconv.ParseInt(tok.literal, 0, 64)
		fmt.Println("value", value)
	} else if tok.kind == "LPAREN" {
		idx++
		value = parseExpression(LOWEST)
		tok := tokens[idx+1]
		if tok.kind != "RPAREN" {
			log.Println("syntax error, no right parentheses")
			value = 0
		}
		idx++
	}

	for idx < len(tokens)-1 && precedence < getPrecedence(1) {
		idx++
		value = parseInfixExpression(value)
	}
	return value
}

func parseInfixExpression(left int64) int64 {
	tok := tokens[idx]
	idx++
	var sum int64
	if tok.kind == "ASTERISK" {
		sum = left * parseExpression(getPrecedence(-1))
	} else if tok.kind == "PLUS" {
		sum = left + parseExpression(getPrecedence(-1))
	} else {
		sum = left
	}

	fmt.Println("operation", sum)
	return sum
}

func getPrecedence(i int) int {
	nextTok := tokens[idx+i]
	nextPrec, ok := precedences[nextTok.kind]
	if ok {
		// log.Println("peekPrece", nextPrec)
		return nextPrec
	}
	// log.Println("lowest", LOWEST)
	return LOWEST
}
