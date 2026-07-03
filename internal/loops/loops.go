package loops

import "strconv"

// SumTo возвращает сумму чисел от 1 до n включительно.
// Если n <= 0, возвращается 0.
func SumTo(n int) int {
	if n <= 0 {
		return 0
	}

	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum
}

// Factorial возвращает n!.
// Для отрицательных n возвращается 0, а для 0 — 1.
func Factorial(n int) int {
	if n < 0 {
		return 0
	}

	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// CountEven считает количество чётных чисел в слайсе.
func CountEven(numbers []int) int {
	count := 0
	for _, number := range numbers {
		if number%2 == 0 {
			count++
		}
	}
	return count
}

// FindFirstNegative возвращает индекс первого отрицательного числа и true.
// Если отрицательных чисел нет, возвращает -1 и false.
func FindFirstNegative(numbers []int) (int, bool) {
	for index, number := range numbers {
		if number < 0 {
			return index, true
		}
	}
	return -1, false
}

// SkipMultiplesOfThree возвращает числа от 1 до n, кроме кратных трём.
func SkipMultiplesOfThree(n int) []int {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		if i%3 == 0 {
			continue
		}
		result = append(result, i)
	}
	return result
}

// MultiplicationRow возвращает строку таблицы умножения для числа n.
func MultiplicationRow(n, limit int) []int {
	if limit <= 0 {
		return []int{}
	}

	row := make([]int, 0, limit)
	for i := 1; i <= limit; i++ {
		row = append(row, n*i)
	}
	return row
}

// ReverseStringByRunes переворачивает строку по рунам, а не по байтам.
func ReverseStringByRunes(s string) string {
	runes := []rune(s)
	for left, right := 0, len(runes)-1; left < right; left, right = left+1, right-1 {
		runes[left], runes[right] = runes[right], runes[left]
	}
	return string(runes)
}

// CountRunes считает, сколько раз каждая руна встречается в строке.
func CountRunes(s string) map[rune]int {
	result := make(map[rune]int)
	for _, r := range s {
		result[r]++
	}
	return result
}

// FizzBuzz возвращает классический FizzBuzz для чисел от 1 до n.
func FizzBuzz(n int) []string {
	result := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%15 == 0:
			result = append(result, "FizzBuzz")
		case i%3 == 0:
			result = append(result, "Fizz")
		case i%5 == 0:
			result = append(result, "Buzz")
		default:
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

// TriangleRows строит текстовый треугольник из символов *.
func TriangleRows(height int) []string {
	rows := make([]string, 0, height)
	for row := 1; row <= height; row++ {
		line := ""
		for col := 0; col < row; col++ {
			line += "*"
		}
		rows = append(rows, line)
	}
	return rows
}

// SumUntilLimit складывает числа, пока сумма не превысит limit.
func SumUntilLimit(numbers []int, limit int) int {
	sum := 0
	for _, number := range numbers {
		if sum+number > limit {
			break
		}
		sum += number
	}
	return sum
}

// Flatten разворачивает матрицу в один слайс.
func Flatten(matrix [][]int) []int {
	result := make([]int, 0)
	for _, row := range matrix {
		for _, value := range row {
			result = append(result, value)
		}
	}
	return result
}

// MaxInSlice возвращает максимальный элемент и true.
// Для пустого слайса возвращает 0 и false.
func MaxInSlice(numbers []int) (int, bool) {
	if len(numbers) == 0 {
		return 0, false
	}

	maxValue := numbers[0]
	for _, number := range numbers[1:] {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue, true
}

// UniquePreserveOrder оставляет только первые появления значений.
func UniquePreserveOrder(numbers []int) []int {
	seen := make(map[int]bool)
	result := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if seen[number] {
			continue
		}
		seen[number] = true
		result = append(result, number)
	}
	return result
}

// RepeatString повторяет строку times раз.
func RepeatString(word string, times int) string {
	if times <= 0 {
		return ""
	}

	result := ""
	for i := 0; i < times; i++ {
		result += word
	}
	return result
}
