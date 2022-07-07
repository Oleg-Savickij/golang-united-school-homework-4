package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	trimmedInput := strings.TrimSpace(input)
	if trimmedInput == "" {
		return "", fmt.Errorf(errorEmptyInput.Error())
	}

	var operandsSlice = strings.Split(trimmedInput, "+")
	if len(operandsSlice) != 2 {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	firstOperand, ok := GetOperand(operandsSlice[0])
	if !ok {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	secondOperand, ok := GetOperand(operandsSlice[1])
	if !ok {
		return "", fmt.Errorf(errorNotTwoOperands.Error())
	}

	result := firstOperand + secondOperand
	return strconv.FormatInt(result, 10), nil
}

func GetOperand(input string) (output int64, ok bool) {
	if input == "" {
		return 0, false
	}
	value, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		return 0, false
	}

	return value, true
}
