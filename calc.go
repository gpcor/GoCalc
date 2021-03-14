package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	calcArguments := os.Args

	firstNum, err := strconv.ParseFloat(calcArguments[1], 64)
	if err != nil {
		log.Fatal("Error:", err)
	}

	secondNum, err := strconv.ParseFloat(calcArguments[2], 64)
	if err != nil {
		log.Fatal("Error:", err)
	}

	operation := calcArguments[3]

	if strings.Contains(operation, "add") {
		fmt.Println(addition(firstNum, secondNum))
	}

	if strings.Contains(operation, "subtract") {
		fmt.Println(subtraction(firstNum, secondNum))
	}
	if strings.Contains(operation, "multiply") {
		fmt.Println(multiplication(firstNum, secondNum))
	}
	if strings.Contains(operation, "divide") {
		fmt.Println(division(firstNum, secondNum))
	}
}
