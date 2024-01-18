package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calculate(operand1, operand2, operator string) (interface{}, error) {
	num1, err := convertToNumber(operand1)
	if err != nil {
		return nil, err
	}

	num2, err := convertToNumber(operand2)
	if err != nil {
		return nil, err
	}

	if (isRomanNumeral(operand1) && num1 > 10) || (isRomanNumeral(operand2) && num2 > 10) {
		return nil, fmt.Errorf("один из операндов больше 10")
	}

	var result int
	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			return nil, fmt.Errorf("деление на ноль")
		}
		result = num1 / num2
	default:
		return nil, fmt.Errorf("неправильный оператор: %s", operator)
	}

	if isRomanNumeral(operand1) && isRomanNumeral(operand2) {
		return convertToRoman(result), nil
	}

	return result, nil
}

func convertToNumber(operand string) (int, error) {
	num, err := strconv.Atoi(operand)
	if err == nil {
		return num, nil
	}

	num, err = convertRomanToArabic(operand)
	if err != nil {
		return 0, fmt.Errorf("неверный формат цифры: %s", operand)
	}

	if num > 10 {
		return 0, fmt.Errorf("операнд больше 10")
	}

	return num, nil
}

func convertToString(num interface{}) string {
	switch v := num.(type) {
	case int:
		return strconv.Itoa(v)
	case string:
		return v
	default:
		return ""
	}
}

func convertRomanToArabic(roman string) (int, error) {
	romanNumerals := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	var result int
	var prevValue int

	for i := len(roman) - 1; i >= 0; i-- {
		value := romanNumerals[rune(roman[i])]

		if value < prevValue {
			result -= value
		} else {
			result += value
		}

		prevValue = value
	}

	return result, nil
}

func convertToRoman(num int) string {
	if num <= 0 || num > 3999 {
		return ""
	}

	romanNumerals := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	var result strings.Builder

	for _, numeral := range romanNumerals {
		for num >= numeral.Value {
			result.WriteString(numeral.Symbol)
			num -= numeral.Value
		}
	}

	return result.String()
}

func isRomanNumeral(s string) bool {
	if len(s) < 1 {
		return false
	}

	romanNumerals := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	for i := 0; i < len(s); i++ {
		if value, ok := romanNumerals[s[i]]; !ok {
			return false
		} else if i > 0 && romanNumerals[s[i-1]] < value {
			return false
		}
	}

	return true
}

func parseInput(input string) ([]string, string, error) {
	var operands []string
	var operator string

	parts := strings.Fields(input)
	if len(parts) != 3 {
		return nil, "", fmt.Errorf("неверное количество элементов в выражении")
	}

	operands = append(operands, parts[0], parts[2])
	operator = parts[1]

	return operands, operator, nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Калькулятор для Kata Academy")
		fmt.Println("Введите выражение: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		operands, operator, err := parseInput(text)
		if err != nil {
			fmt.Println(err)
			continue
		}

		result, err := calculate(operands[0], operands[1], operator)
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("Результат: ", convertToString(result))
	}
}
