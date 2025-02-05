package main

import (
	"fmt"
	maths "simple_calculator/modules"
	"log"
)

func main() {
	var operator string
	fmt.Println("What operation do you want to do? (+, -, *, /)")
	fmt.Scan(&operator)

	switch operator {
	// addition case
	case "+":
		a, b, err := askValues()
		if err != nil {
			log.Fatal("Error reading numbers")
		}
		var result int = maths.DoAdd(a, b)
		fmt.Println("---")
		fmt.Printf("The sum of %v & %v is %v\n", a, b, result)
	// subtraction case
	case "-":
		a, b, err := askValues()
		if err != nil {
			log.Fatal("Error reading numbers")
		}
		var result int = maths.DoSubtract(a, b)
		fmt.Println("---")
		fmt.Printf("The difference of %v & %v is %v\n", a, b, result)
	// multiplication case
	case "*":
		a, b, err := askValues()
		if err != nil {
			log.Fatal("Error reading numbers")
		}
		var result int = maths.DoMultiply(a, b)
		fmt.Println("---")
		fmt.Printf("The product of %v & %v is %v\n", a, b, result)
	// division case
	case "/":
		a, b, err := askValues()
		if err != nil {
			log.Fatal("Error reading numbers")
		}
		result, remainder, err := maths.DoDivide(a, b)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println("---")
		fmt.Printf("The quotient of %v & %v is %v and remainder is %v\n", a, b, result, remainder)
	default:
		fmt.Println("Please enter a valid operator.")
	}
}

func askValues() (int, int, error) {
	var a int
	fmt.Println("What is the first number?")
	_, err := fmt.Scan(&a)
	if err != nil {
		log.Fatal("Error reading number")
	}

	var b int
	fmt.Println("What is the second number?")
	_, err = fmt.Scan(&b)
	if err != nil {
		log.Fatal("Error reading number")
	}

	return a, b, err
}
