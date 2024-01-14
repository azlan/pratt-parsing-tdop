package main

type token struct {
	kind    string
	literal string
}

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

var (
	idx    = 0
	tokens = []token{}
)

// naive tokenizer just get job done
func getTokens(input string) {
	idx = 0
	tokens = []token{}
	for i := 0; i < len(input); i++ {
		ch := input[i]
		if isWhitespace(ch) {
			continue
		}
		if isDigit(ch) {
			number := string(ch)
			i++
			if i == len(input) {
				tokens = append(tokens, token{
					kind: "NUMBER", literal: number,
				})
				break
			}
			for isDigit(input[i]) {
				number += string(input[i])
				i++
			}
			i--
			tokens = append(tokens, token{
				kind: "NUMBER", literal: number,
			})
		}
		if ch == '+' {
			tokens = append(tokens, token{kind: "PLUS", literal: "+"})
		}
		if ch == '*' {
			tokens = append(tokens, token{kind: "ASTERISK", literal: "*"})
		}
		if ch == '(' {
			tokens = append(tokens, token{kind: "LPAREN", literal: "("})
		}
		if ch == ')' {
			tokens = append(tokens, token{kind: "RPAREN", literal: ")"})
		}
	}
}

// func isLetter(ch byte) bool {
// 	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
// }

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func isWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}
