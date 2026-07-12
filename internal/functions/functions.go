package functions

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func SafeDivide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func FindActiveUser(users []User, id int) (User, bool) {
	for _, user := range users {
		if user.ID == id && user.Active {
			return user, true
		}
	}
	return User{}, false
}

func SumAll(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Apply(a, b int, op func(int, int) int) int {
	return op(a, b)
}

func NewCounter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func DeferOrder() (result string) {
	// Тема: несколько defer выполняются в обратном порядке — LIFO.
	// Важно: defer выполнится после того, как return-значение уже подготовлено,
	// но до фактического выхода из функции.
	result = "first" // "first"
	defer func() { result += "-third" }()
	defer func() { result += "-second" }()
	return
}

func DeferArguments() (result string) {
	// Тема: аргументы defer вычисляются сразу, в момент объявления defer.
	// Частая ошибка: думать, что defer возьмёт новое значение переменной при выходе из функции.
	value := "first" // "first"
	defer func(captured string) {
		result = "captured=" + captured + " current=" + value
	}(value)

	value = "second" // "second"
	result = "body"
	return
}

func Example() string {
	var out strings.Builder

	// Тема: функция — маленькое именованное действие с понятной сигнатурой.
	// В этом блоке уже можно использовать всё, что прошли раньше: errors, structs, slices, loops.
	quotient, err := SafeDivide(10, 2) // 5, nil
	errorText := "nil"                 // "nil"
	if err != nil {
		errorText = err.Error()
	}
	fmt.Fprintf(&out, "divide=%d\n", quotient)
	fmt.Fprintf(&out, "err=%s\n", errorText)

	users := []User{{ID: 1, Name: "Maria", Active: true}, {ID: 2, Name: "Alex", Active: false}} // [{1 Maria true} {2 Alex false}]
	user, ok := FindActiveUser(users, 1)                                                        // {1 Maria true}, true
	fmt.Fprintf(&out, "user=%s\n", user.Name)
	fmt.Fprintf(&out, "ok=%t\n", ok)

	sum := SumAll(1, 2, 3, 4) // 10
	fmt.Fprintf(&out, "sum=%d\n", sum)

	product := Apply(3, 4, func(a, b int) int { return a * b }) // 12
	fmt.Fprintf(&out, "product=%d\n", product)

	counter := NewCounter()
	first := counter()  // 1
	second := counter() // 2
	fmt.Fprintf(&out, "counter=%d/%d\n", first, second)

	// Тема: defer выполняется при выходе из функции в обратном порядке.
	deferOrder := DeferOrder() // first-second-third
	fmt.Fprintf(&out, "defer=%s\n", deferOrder)

	// Тема: параметры deferred-вызова вычисляются сразу.
	// captured останется "first", хотя value позже изменился на "second".
	deferArguments := DeferArguments() // captured=first current=second
	fmt.Fprintf(&out, "args=%s", deferArguments)

	// Ответ: увидим результат деления, найденного пользователя, variadic, function value, closure и defer.
	return out.String()
}
