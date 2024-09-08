package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Println("Введите операцию")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		operation := strings.Split(text, " ")
		if len(operation) <= 2 {
			err := errors.New("выдача паники, так как строка не является математической операцией")
			fmt.Println(err)
			break
		}
		if len(operation) > 3 {
			err := errors.New("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
			fmt.Println(err)
			break
		}
		firstNum := operation[0]
		secondNum := operation[2]
		sign := operation[1]
		switch sign {
		case "+":
			fmt.Println(sum(firstNum, secondNum))
			break loop
		case "-":
			fmt.Println(sub(firstNum, secondNum))
			break loop
		case "*":
			fmt.Println(multiply(firstNum, secondNum))
			break loop
		case "/":
			fmt.Println(divide(firstNum, secondNum))
			break loop
		default:
			err := errors.New("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
			fmt.Println(err)
			break loop
		}
	}
}

func detector(number string) bool {
	for _, char := range number {
		if !strings.ContainsRune("IVX", char) {
			return false
		}
	}
	return true
}

func romeToArab(number string) int {
	romanNumerals := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
	}
	total := 0
	prevValue := 0

	for _, char := range number {
		value := romanNumerals[string(char)]
		if value > prevValue {
			total += value - 2*prevValue
		} else {
			total += value
		}
		prevValue = value
	}
	return total
}

func arabToRome(number int) string {
	romanNumerals := []struct {
		value  int
		symbol string
	}{
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
	result := ""

	for _, rn := range romanNumerals {
		for number >= rn.value {
			result += rn.symbol
			number -= rn.value
		}
	}
	return result
}

func sum(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) < 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := romeToArab(firstNum) + romeToArab(secondNum)
		return arabToRome(result)
	} else if detector(firstNum) == false && detector(secondNum) == false {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := firstNumInt + secondNumInt
		return result
	}
	err := errors.New("выдача паники, так как используются одновременно разные системы счисления")
	return err
}

func sub(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) < 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := romeToArab(firstNum) - romeToArab(secondNum)
		if result <= 0 {
			err := errors.New("выдача паники, так как в римской системе нет отрицательных чисел")
			return err
		}
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := firstNumInt - secondNumInt
		return result
	}
	err := errors.New("выдача паники, так как используются одновременно разные системы счисления")
	return err
}

func multiply(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) < 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := romeToArab(firstNum) * romeToArab(secondNum)
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := firstNumInt * secondNumInt
		return result
	}
	err := errors.New("выдача паники, так как используются одновременно разные системы счисления")
	return err
}

func divide(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) < 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := romeToArab(firstNum) / romeToArab(secondNum)
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			err := errors.New("выдача паники, число не должно быть < 1 или > 10")
			return err
		}
		result := firstNumInt / secondNumInt
		return result
	}
	err := errors.New("выдача паники, так как используются одновременно разные системы счисления")
	return err
}