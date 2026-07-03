package arrays

// ZeroArray возвращает массив из трёх int со zero value.
func ZeroArray() [3]int {
	return [3]int{}
}

// Weekdays возвращает массив дней недели.
func Weekdays() [7]string {
	return [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
}

// SetAt возвращает копию массива с изменённым элементом.
func SetAt(numbers [5]int, index int, value int) [5]int {
	if index < 0 || index >= len(numbers) {
		return numbers
	}
	numbers[index] = value
	return numbers
}

// Sum считает сумму элементов массива.
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Average считает среднее значение массива.
func Average(numbers [5]int) float64 {
	return float64(Sum(numbers)) / float64(len(numbers))
}

// Max возвращает максимальный элемент массива.
func Max(numbers [5]int) int {
	maxValue := numbers[0]
	for _, number := range numbers[1:] {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue
}

// Reverse возвращает новый массив в обратном порядке.
func Reverse(numbers [5]int) [5]int {
	var result [5]int
	for i := range numbers {
		result[len(numbers)-1-i] = numbers[i]
	}
	return result
}

// Contains проверяет наличие значения в массиве.
func Contains(numbers [5]int, target int) bool {
	for _, number := range numbers {
		if number == target {
			return true
		}
	}
	return false
}

// CountValue считает количество вхождений значения.
func CountValue(numbers [5]int, target int) int {
	count := 0
	for _, number := range numbers {
		if number == target {
			count++
		}
	}
	return count
}

// Equal сравнивает два массива.
func Equal(left, right [5]int) bool {
	return left == right
}

// CopyAndSet демонстрирует value semantics: массив копируется.
func CopyAndSet(numbers [5]int, index int, value int) ([5]int, [5]int) {
	copyOfNumbers := numbers
	if index >= 0 && index < len(copyOfNumbers) {
		copyOfNumbers[index] = value
	}
	return numbers, copyOfNumbers
}

// FirstLast возвращает первый и последний элемент.
func FirstLast(numbers [5]int) (int, int) {
	return numbers[0], numbers[len(numbers)-1]
}

// ToSlice возвращает слайс-копию массива.
func ToSlice(numbers [5]int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers[:])
	return result
}

// MatrixDiagonalSum считает сумму главной диагонали 3x3.
func MatrixDiagonalSum(matrix [3][3]int) int {
	sum := 0
	for i := 0; i < len(matrix); i++ {
		sum += matrix[i][i]
	}
	return sum
}

// CompareBySum сравнивает массивы по сумме элементов.
func CompareBySum(left, right [5]int) string {
	leftSum := Sum(left)
	rightSum := Sum(right)
	switch {
	case leftSum > rightSum:
		return "left"
	case leftSum < rightSum:
		return "right"
	default:
		return "equal"
	}
}
