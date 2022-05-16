package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
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
	var first, second int
	input = NormalString(input)
	if input == "" {
		return "", fmt.Errorf("input string is empty: %w", errorEmptyInput)
	}
	splited := strings.FieldsFunc(input,
		func(r rune) bool {
			return strings.ContainsRune("+-", r)
		})
	if len(splited) != 2 {
		return "", fmt.Errorf("two operands not found: %w",errorNotTwoOperands)
	}
	if strings.HasPrefix(input, "-") {
		first, _ = strconv.Atoi(splited[0])
		first = -1 * first
	} else {
		first, _ = strconv.Atoi(splited[0])
	}
	if strings.LastIndex(input, "-") > 0 {
		second, _ = strconv.Atoi(splited[1])
		second = -1 * second
	} else {
		second, _ = strconv.Atoi(splited[1])
	}
	return fmt.Sprint(first + second), nil
}

func NormalString(input string) string {
	var normal []rune
	for _, r := range input {
		if unicode.IsSpace(r) {
			continue
		}
		normal = append(normal, r)
	}
	return string(normal)
}
