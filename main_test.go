package main

import (
	"fmt"
	"testing"
)

func TestString1(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"2+1* 3", 5},
		{" 1 + 1 ", 2},
		{"2 * 3 ", 6},
		{" 2  + 1 * 3 ", 5},
		{"2 + 3*3 + 1", 12},
	}

	for i, tt := range tests {
		getTokens(tt.input)
		sum1 := parseExpression(LOWEST)
		fmt.Printf("result sum: %d\n\n", sum1)

		if sum1 != tt.expected {
			t.Logf("test[%d] - expected=%d, got=%d", i, tt.expected, sum1)
			t.Fail()
		}

	}
}

func TestString2(t *testing.T) {
	input := "1 + 2"
	getTokens(input)
	sum1 := parseExpression(LOWEST)
	fmt.Println("result sum:", sum1)
	expected := int64(3)
	if sum1 != expected {
		t.Fatalf("expected=%d, got=%d", expected, sum1)

	}
}

func TestParens(t *testing.T) {
	input := "(1 + 2) * 2"
	getTokens(input)
	sum1 := parseExpression(LOWEST)
	fmt.Println("result sum:", sum1)
	expected := int64(6)
	if sum1 != expected {
		t.Fatalf("expected=%d, got=%d", expected, sum1)

	}
}

func TestParens2(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"(2 + 3 * (3 +1))", 14},
		{"(2 + 3) + (3 * 2)", 11},
	}

	for i, tt := range tests {
		getTokens(tt.input)
		sum1 := parseExpression(LOWEST)
		fmt.Printf("result sum: %d\n\n", sum1)

		if sum1 != tt.expected {
			t.Logf("test[%d] - expected=%d, got=%d", i, tt.expected, sum1)
			t.Fail()
		}

	}
}
