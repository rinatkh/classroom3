package slices

import "sort"

// NewSlice создаёт слайс из переданных чисел.
func NewSlice(numbers ...int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	return result
}

// AppendValue возвращает новый слайс с добавленным значением.
func AppendValue(numbers []int, value int) []int {
	return append(numbers, value)
}

// Sum считает сумму элементов слайса.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Average считает среднее значение.
func Average(numbers []int) (float64, bool) {
	if len(numbers) == 0 {
		return 0, false
	}
	return float64(Sum(numbers)) / float64(len(numbers)), true
}

// FilterEven возвращает только чётные числа.
func FilterEven(numbers []int) []int {
	result := make([]int, 0)
	for _, number := range numbers {
		if number%2 == 0 {
			result = append(result, number)
		}
	}
	return result
}

// MapDouble умножает каждый элемент на 2.
func MapDouble(numbers []int) []int {
	result := make([]int, len(numbers))
	for i, number := range numbers {
		result[i] = number * 2
	}
	return result
}

// FindIndex возвращает индекс значения или -1.
func FindIndex(numbers []int, target int) int {
	for i, number := range numbers {
		if number == target {
			return i
		}
	}
	return -1
}

// RemoveAt удаляет элемент по индексу и возвращает новый слайс.
func RemoveAt(numbers []int, index int) []int {
	if index < 0 || index >= len(numbers) {
		return append([]int(nil), numbers...)
	}
	result := make([]int, 0, len(numbers)-1)
	result = append(result, numbers[:index]...)
	result = append(result, numbers[index+1:]...)
	return result
}

// InsertAt вставляет значение по индексу.
func InsertAt(numbers []int, index int, value int) []int {
	if index < 0 {
		index = 0
	}
	if index > len(numbers) {
		index = len(numbers)
	}
	result := make([]int, 0, len(numbers)+1)
	result = append(result, numbers[:index]...)
	result = append(result, value)
	result = append(result, numbers[index:]...)
	return result
}

// CopySlice возвращает независимую копию слайса.
func CopySlice(numbers []int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	return result
}

// ReverseInPlace переворачивает слайс на месте.
func ReverseInPlace(numbers []int) {
	for left, right := 0, len(numbers)-1; left < right; left, right = left+1, right-1 {
		numbers[left], numbers[right] = numbers[right], numbers[left]
	}
}

// Unique возвращает уникальные значения с сохранением порядка.
func Unique(numbers []int) []int {
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

// Window возвращает под-слайс безопасно.
func Window(numbers []int, start, end int) []int {
	if start < 0 {
		start = 0
	}
	if end > len(numbers) {
		end = len(numbers)
	}
	if start > end {
		return []int{}
	}
	return append([]int(nil), numbers[start:end]...)
}

// Chunk делит слайс на части указанного размера.
func Chunk(numbers []int, size int) [][]int {
	if size <= 0 {
		return [][]int{}
	}
	result := make([][]int, 0)
	for start := 0; start < len(numbers); start += size {
		end := start + size
		if end > len(numbers) {
			end = len(numbers)
		}
		result = append(result, append([]int(nil), numbers[start:end]...))
	}
	return result
}

// MergeAndSort объединяет два слайса и сортирует результат.
func MergeAndSort(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	result = append(result, left...)
	result = append(result, right...)
	sort.Ints(result)
	return result
}
