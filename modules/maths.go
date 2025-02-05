package maths

import "errors"

func DoAdd(a int, b int) int {
	var result int = a + b
	return result
}

func DoSubtract(a int, b int) int {
	var result int = a - b
	return result
}

func DoMultiply(a int, b int) int {
	var result int = a * b
	return result
}

func DoDivide(a int, b int) (int, int, error) {
	var err error
	if b == 0 {
		err = errors.New("dividing by 0 not allowed")
		return 0, 0, err
	}
	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
