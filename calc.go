package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	calcArguments := os.Args

	firstNum, err := strconv.ParseFloat(calcArguments[1], 64)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	secondNum, err := strconv.ParseFloat(calcArguments[2], 64)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	switch operation := calcArguments[3]; operation {
	case "add":
		fmt.Println(addition(firstNum, secondNum))
	case "subtract":
		fmt.Println(subtraction(firstNum, secondNum))
	case "multiply":
		fmt.Println(multiplication(firstNum, secondNum))
	case "divide":
		fmt.Println(division(firstNum, secondNum))
	}
}
