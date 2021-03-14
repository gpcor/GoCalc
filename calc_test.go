package main

import "testing"

func benchmarkCalculations(b *testing.B) {
	firstNumber := 12.1
	secondNumber := 12.1

	addition(firstNumber, secondNumber)
	subtraction(firstNumber, secondNumber)
	multiplication(firstNumber, secondNumber)
	division(firstNumber, secondNumber)
}
