package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
		romanResult, err := convertToRoman(result)
		if err != nil {
			return nil, err
		}
		return romanResult, nil
	}

	return result, nil
}

func convertToNumber(operand string) (int, error) {
	for _, char := range operand {
		if !unicode.IsDigit(char) && char != 'I' && char != 'V' && char != 'X' {
			return 0, fmt.Errorf("неверный формат цифры: %s", operand)
		}
	}

	if isRomanNumeral(operand) {
		return convertRomanToArabic(operand)
	}

	num, err := strconv.Atoi(operand)
	if err != nil {
		return 0, fmt.Errorf("неверный формат цифры: %s", operand)
	}

	if num < 1 || num > 10 {
		return 0, fmt.Errorf("операнд должен быть в пределах от 1 до 10")
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

func convertToRoman(num int) (string, error) {
	if num <= 0 || num > 3999 {
		return "", fmt.Errorf("в римской числовой системе нет отрицательных чисел или чисел больше 3999")
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

	return result.String(), nil
}

func isRomanNumeral(s string) bool {
	if len(s) < 1 {
		return false
	}

	for _, char := range s {
		if char != 'I' && char != 'V' && char != 'X' {
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

	operand1, operand2 := parts[0], parts[2]

	if (isRomanNumeral(operand1) && !isRomanNumeral(operand2)) || (!isRomanNumeral(operand1) && isRomanNumeral(operand2)) {
		return nil, "", fmt.Errorf("используйте либо римские, либо арабские числа в одном выражении")
	}

	operands = append(operands, operand1, operand2)
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
