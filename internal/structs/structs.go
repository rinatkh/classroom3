package structs

import (
	"fmt"
	"strings"
)

type User struct {
	ID     int
	Name   string
	Active bool
}

func Example() string {
	var out strings.Builder

	// Тема: структура группирует связанные поля в один тип.
	// Пока без методов: только данные и доступ к полям через точку.
	user := User{ID: 1, Name: "Maria", Active: true} // {1 Maria true}

	// Тема: поле можно читать и менять через точку.
	// Частая ошибка: забыть имя поля и пытаться обращаться как к map.
	before := user.Name // Maria
	user.Name = "Masha"

	fmt.Fprintf(&out, "before=%s\n", before)
	fmt.Fprintf(&out, "after=%s\n", user.Name)

	// Тема: zero value структуры — это zero value каждого поля.
	var empty User // {0 "" false}
	fmt.Fprintf(&out, "empty=%+v\n", empty)

	fmt.Fprintf(&out, "active=%t", user.Active)

	// Ответ: увидим старое имя, новое имя и zero value структуры.
	return out.String()
}
