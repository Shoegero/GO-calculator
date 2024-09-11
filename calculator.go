package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Введите операцию")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		operation := strings.Split(text, " ")
		if len(operation) <= 2 {
			panic("выдача паники, так как строка не является математической операцией")
		}
		if len(operation) > 3 {
			panic("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		}
		firstNum := operation[0]
		secondNum := operation[2]
		sign := operation[1]
		switch sign {
		case "+":
			fmt.Println(sum(firstNum, secondNum))
		case "-":
			fmt.Println(sub(firstNum, secondNum))
		case "*":
			fmt.Println(multiply(firstNum, secondNum))
		case "/":
			fmt.Println(divide(firstNum, secondNum))
		default:
			panic("выдача паники, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)")
		}
	}
}

func detector(number string) bool {
	romanNumbers := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabicNumbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for i := 0; i < len(romanNumbers); i++ {
		if romanNumbers[i] == number {
			return true
		}
	}
	for i := 0; i < len(arabicNumbers); i++ {
		if arabicNumbers[i] == number {
			return false
		}
	}
	panic("выдача паники, так как строка не является математической операцией")
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
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := romeToArab(firstNum) + romeToArab(secondNum)
		return arabToRome(result)
	} else if detector(firstNum) == false && detector(secondNum) == false {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := firstNumInt + secondNumInt
		return result
	}
	panic("выдача паники, так как используются одновременно разные системы счисления")
}

func sub(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := romeToArab(firstNum) - romeToArab(secondNum)
		if result <= 0 {
			panic("выдача паники, так как в римской системе нет отрицательных чисел и 0")
		}
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := firstNumInt - secondNumInt
		return result
	}
	panic("выдача паники, так как используются одновременно разные системы счисления")
}

func multiply(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := romeToArab(firstNum) * romeToArab(secondNum)
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := firstNumInt * secondNumInt
		return result
	}
	panic("выдача паники, так как используются одновременно разные системы счисления")
}

func divide(firstNum string, secondNum string) interface{} {
	if detector(firstNum) && detector(secondNum) {
		if romeToArab(firstNum) < 1 || romeToArab(firstNum) > 10 || romeToArab(secondNum) < 1 || romeToArab(secondNum) > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := romeToArab(firstNum) / romeToArab(secondNum)
		if result == 0 {
			panic("выдача паники, так как в римской системе нет числа 0")
		}
		return arabToRome(result)
	} else if !(detector(firstNum) && detector(secondNum)) {
		firstNumInt, _ := strconv.Atoi(firstNum)
		secondNumInt, _ := strconv.Atoi(secondNum)
		if firstNumInt < 1 || firstNumInt > 10 || secondNumInt < 1 || secondNumInt > 10 {
			panic("выдача паники, число не должно быть < 1 или > 10")
		}
		result := firstNumInt / secondNumInt
		return result
	}
	panic("выдача паники, так как используются одновременно разные системы счисления")
}
