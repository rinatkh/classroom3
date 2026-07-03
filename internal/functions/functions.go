package functions

import (
	"fmt"
	"strings"
)

// FullName склеивает имя и фамилию через пробел.
func FullName(firstName, lastName string) string {
	return strings.TrimSpace(firstName + " " + lastName)
}

// PriceWithDiscount применяет скидку в процентах.
func PriceWithDiscount(price float64, discountPercent float64) float64 {
	if price < 0 || discountPercent < 0 {
		return 0
	}
	if discountPercent > 100 {
		discountPercent = 100
	}
	return price * (100 - discountPercent) / 100
}

// SafeDivide делит a на b и возвращает флаг успеха.
func SafeDivide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a / b, true
}

// MinMax возвращает минимальное и максимальное значение.
func MinMax(a, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// NormalizeEmail приводит email к нижнему регистру и убирает пробелы.
func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

// FormatUser формирует короткую строку пользователя.
func FormatUser(id int, name string, active bool) string {
	status := "inactive"
	if active {
		status = "active"
	}
	return fmt.Sprintf("#%d %s (%s)", id, name, status)
}

// ApplyOperation применяет переданную функцию к двум числам.
func ApplyOperation(a, b int, operation func(int, int) int) int {
	if operation == nil {
		return 0
	}
	return operation(a, b)
}

// SumVariadic складывает любое количество чисел.
func SumVariadic(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// BuildGreeting возвращает приветствие на выбранном языке.
func BuildGreeting(language, name string) string {
	switch language {
	case "ru":
		return "Привет, " + name
	case "en":
		return "Hello, " + name
	default:
		return "Hi, " + name
	}
}

// ValidatePassword проверяет пароль и возвращает объяснение.
func ValidatePassword(password string) (bool, string) {
	if len(password) < 8 {
		return false, "password is too short"
	}
	if !strings.ContainsAny(password, "0123456789") {
		return false, "password must contain digit"
	}
	return true, "ok"
}

// SplitFullName разделяет строку на имя и фамилию.
func SplitFullName(fullName string) (string, string) {
	parts := strings.Fields(fullName)
	if len(parts) == 0 {
		return "", ""
	}
	if len(parts) == 1 {
		return parts[0], ""
	}
	return parts[0], parts[1]
}

// CalcOrderTotal считает итог заказа с доставкой и скидкой.
func CalcOrderTotal(price, delivery, discount float64) float64 {
	return PriceWithDiscount(price, discount) + delivery
}

// MakeCounter возвращает функцию-счётчик.
func MakeCounter(start int) func() int {
	value := start
	return func() int {
		value++
		return value
	}
}

// Swap меняет два значения местами и возвращает результат.
func Swap(a, b string) (string, string) {
	return b, a
}

// Average возвращает среднее значение и флаг успеха.
func Average(numbers ...int) (float64, bool) {
	if len(numbers) == 0 {
		return 0, false
	}
	return float64(SumVariadic(numbers...)) / float64(len(numbers)), true
}
