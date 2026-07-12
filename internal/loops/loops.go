package loops

import (
	"fmt"
	"strings"
)

type User struct {
	Name   string
	Active bool
}

func Example() string {
	var out strings.Builder

	// Тема: классический for состоит из init; condition; post.
	sum := 0 // 0
	for i := 1; i <= 5; i++ {
		sum += i // 1, 3, 6, 10, 15
	}
	fmt.Fprintf(&out, "sum=%d\n", sum)

	// Тема: for condition заменяет while.
	countdown := 3 // 3
	for countdown > 0 {
		countdown-- // 2, 1, 0
	}
	fmt.Fprintf(&out, "countdown=%d\n", countdown)

	// Тема: for { } — бесконечный цикл. Нужен break, return или context.
	iterations := 0 // 0
	for {
		iterations++ // 1, 2
		if iterations == 2 {
			break
		}
	}
	fmt.Fprintf(&out, "iterations=%d\n", iterations)

	// Тема: range удобен для обхода слайсов и массивов.
	// Частая ошибка: менять value, а не numbers[i]. value — копия.
	numbers := []int{1, 2, 3} // [1 2 3]
	for i := range numbers {
		numbers[i] *= 10 // [10 20 30]
	}
	fmt.Fprintf(&out, "numbers=%v\n", numbers)

	// Тема: цикл по слайсу структур похож на обработку данных из БД или API.
	users := []User{{Name: "Maria", Active: true}, {Name: "Alex", Active: false}} // [{Maria true} {Alex false}]
	activeCount := 0                                                              // 0
	for _, user := range users {
		if user.Active {
			activeCount++ // 1
		}
	}
	fmt.Fprintf(&out, "active=%d", activeCount)

	// Ответ: увидим сумму, завершённый while-like цикл, break, изменённый слайс и количество активных пользователей.
	return out.String()
}
