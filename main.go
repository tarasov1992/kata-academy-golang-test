package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Expression struct {
	X, Y     int
	Operator string
	isRoman  bool
}

func arabicToRoman(num int) string {
	romanNum := ""
	arabicNums := []int{1, 4, 5, 9, 10, 40, 50, 90, 100, 400, 500, 900, 1000}
	romanNums := []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C", "CD", "D", "CM", "M"}
	i := len(arabicNums) - 1

	for num > 0 {
		for arabicNums[i] <= num {
			romanNum += romanNums[i]
			num -= arabicNums[i]
		}
		i -= 1
	}

	return romanNum
}

func (exp Expression) Calculate() (arabic int, roman string, err error) {
	var result int

	switch {
	case exp.Operator == "+":
		result = exp.X + exp.Y
	case exp.Operator == "-":
		result = exp.X - exp.Y
	case exp.Operator == "*":
		result = exp.X * exp.Y
	case exp.Operator == "/":
		result = exp.X / exp.Y
	}

	if (result < 1 || result > 3999) && exp.isRoman == true {
		return 0, roman, errors.New("your expression's result can't be validly displayed as a roman numeral")
	} else if exp.isRoman == true {
		exp.isRoman = true
		romanResult := arabicToRoman(result)
		return 0, romanResult, nil
	} else {
		return result, roman, nil
	}
}

func extractNumber(str string) (int, bool, error) {
	var num int
	var isRoman bool
	var e error

	romanNums := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	res, err := strconv.Atoi(str)
	if err != nil && romanNums[str] == 0 {
		// if Atoi can't convert to a number and the part isn't a roman num, return err
		e = errors.New("one of your variables isn't a whole arabic number or roman numeral, or your roman numeral isn't within the I..X range")
	} else if err != nil && romanNums[str] != 0 {
		// if Atoi can't convert to a number but the str is a roman num, assign arabic num to x
		num = romanNums[str]
		isRoman = true
	} else {
		num = res
	}

	if e != nil {
		return 0, false, e
	} else {
		return num, isRoman, nil
	}
}

func validateOperator(str string) (string, error) {
	operators := []string{"+", "-", "*", "/"}
	containsOperator := slices.Contains(operators, str)
	if containsOperator == false {
		return "", errors.New("your operator is invalid, allowed operators are: +, -, *, /")
	} else {
		return str, nil
	}
}

func convertStrIntoExp(usrInput string) (exp Expression, err error) {

	// remove new line token
	expStr := strings.ReplaceAll(usrInput, "\n", "")

	// extract parts
	parts := strings.Split(expStr, " ")

	// if there number of parts !=3, the expression is invalid
	if len(parts) != 3 {
		return exp, errors.New("invalid expression, an expression should have 2 numbers and an operand separated by spaces")
	}

	x, isXroman, Xerr := extractNumber(parts[0])
	if Xerr != nil {
		return exp, Xerr
	}
	y, isYroman, Yerr := extractNumber(parts[2])
	if Yerr != nil {
		return exp, Yerr
	}

	if (x < 1 || x > 10) || (y < 1 || y > 10) {
		// if the parsed num isn't within the required constraints, return err
		return exp, errors.New("one of your variables is out of the valid range: 1..10 or I..X")
	}

	if isXroman != isYroman {
		return exp, errors.New("your expression should use EITHER arabic OR roman numbers")
	}

	operator, opErr := validateOperator(parts[1])
	if opErr != nil {
		return exp, opErr
	}

	return Expression{x, y, operator, isXroman}, nil
}

func main() {
	// declaring a var that will read and store input from stdin
	buffer := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Input:")

		// reading into a buffer
		usrInputStr, _ := buffer.ReadString('\n')

		// converting a string into an expression
		exp, err := convertStrIntoExp(usrInputStr)
		if err != nil {
			fmt.Println(err)
			break
		}

		arabicRes, romanRes, err := exp.Calculate()
		if err != nil {
			fmt.Println("Output:")
			fmt.Println(err)
			break
		} else if exp.isRoman == true {
			fmt.Println("Output:")
			fmt.Println(romanRes)
		} else {
			fmt.Println("Output:")
			fmt.Println(arabicRes)
		}
	}
}
