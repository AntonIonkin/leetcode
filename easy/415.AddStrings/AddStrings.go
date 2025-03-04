package main

import "fmt"

func main() {
	num1 := "109"
	num2 := "11"

	fmt.Println(addStrings(num1, num2))
}

// addStrings складывает два числа в виде строк и возвращает результат как строку.
func addStrings(num1 string, num2 string) string {
	const byteToDigit byte = '0' // Константа для преобразования байта в цифру.

	var remainder int // Переменная для хранения остатка от сложения.

	// Убедимся, что num1 длиннее или равен num2.
	if len(num1) < len(num2) {
		num1, num2 = num2, num1 // Если num1 короче, меняем местами.
	}
	minStrRightIdx := len(num2) - 1 // Индекс правого конца для num2.

	// Создаем срез байтов для хранения результата сложения.
	// Длина будет максимальной из чисел + 1 (на случай переноса).
	res := make([]byte, max(len(num1), len(num2))+1)
	resIndex := len(res) - 1 // Индекс для записи результата.

	// Проходим по первой строке справа налево.
	for i := len(num1) - 1; i >= 0; i-- {
		// Получаем текущее значение для суммы: текущая цифра + остаток.
		curSum := num1[i] - byteToDigit + byte(remainder)
		remainder = 0 // Сбрасываем остаток.

		// Если есть цифры во втором числе, добавляем их к сумме.
		if minStrRightIdx >= 0 {
			curSum += num2[minStrRightIdx] - byteToDigit // Добавляем цифру из num2.
		}

		// Записываем последнюю цифру суммы в результат и вычисляем новый остаток.
		res[resIndex], remainder = curSum%10+byteToDigit, int(curSum/10)

		// Переходим к следующему индексу в num2 и res.
		minStrRightIdx--
		resIndex--
	}

	// Если остался перенос, добавляем его в начало результата.
	if remainder == 1 {
		res[0] = 1 + byteToDigit // Устанавливаем первый элемент как "1".
		return string(res)       // Возвращаем строку результата.
	}

	// Возвращаем результат, начиная с первого не нулевого байта (результат без ведущих нулей).
	return string(res[resIndex+1:])
}

/*
func addStrings(num1 string, num2 string) string {
	minNum := []rune(num1)
	maxNum := []rune(num2)

	if len(num1) > len(num2) {
		maxNum, minNum = minNum, maxNum
	}

	j := 0
	remains := 0
	result := []rune{}

	for i := len(maxNum) - 1; i >= 0; i-- {
		sum := int(maxNum[i] - '0')

		if j < len(minNum) {
			sum += int(minNum[len(minNum)-1-j] - '0')
			j++
		}

		sum += remains

		if sum >= 10 {
			remains = 1
			sum -= 10
		} else {
			remains = 0
		}

		result = append([]rune{rune(sum + '0')}, result...)

	}

	if remains > 0 {
		result = append([]rune{'1'}, result...)
	}

	return string(result)
}
*/
