package errorsdemo

import (
	"errors"
	"fmt"
	"strings"
)

var ErrNotFound = errors.New("not found")

type FieldError struct {
	Field  string
	Reason string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Reason)
}

func Example() string {
	var out strings.Builder

	// Тема: error в Go — это обычное значение.
	// nil означает, что ошибки нет.
	// Частая ошибка: думать, что error — это что-то отдельное от переменной.
	// На самом деле error можно сравнивать с nil, передавать в функции и возвращать как обычное значение.
	var err error
	fmt.Fprintf(&out, "1) nil error -> %t\n", err == nil) // true

	// Тема: errors.New создаёт простую ошибку с текстом.
	// Это удобно, когда нам не нужен свой тип ошибки и не нужны дополнительные поля.
	// Частая ошибка: игнорировать err и продолжать бизнес-логику как будто всё хорошо.
	err = errors.New("email is empty")
	fmt.Fprintf(&out, "2) errors.New -> %s\n", err) // email is empty

	// Тема: fmt.Errorf с %w оборачивает ошибку
	// и сохраняет исходную причину внутри цепочки ошибок.
	// Это важно, если потом мы хотим использовать errors.Is.
	// Частая ошибка: использовать %s или %v, а потом удивляться, почему errors.Is не работает.
	wrappedNotFound := fmt.Errorf("load user: %w", ErrNotFound)
	fmt.Fprintf(&out, "3) errors.Is with %%w -> %t\n", errors.Is(wrappedNotFound, ErrNotFound)) // true

	// Специально неправильный пример:
	// Так делать НЕ нужно, если потом мы хотим проверять ошибку через errors.Is.
	// %v вставляет только текст ошибки, но НЕ сохраняет саму ошибку в цепочке.
	// Визуально текст будет похожий, но errors.Is уже не сработает.
	wrappedNotFoundWrong := fmt.Errorf("load user: %v", ErrNotFound)
	fmt.Fprintf(&out, "4) anti-example with %%v -> %t\n", errors.Is(wrappedNotFoundWrong, ErrNotFound)) // false

	// Тема: errors.Is обычно используют внутри if,
	// чтобы выбрать разное поведение для разных ошибок.
	// Например: если пользователь не найден — возвращаем одну ветку логики,
	// если ошибка другая — обрабатываем иначе.
	branch := "branch: unexpected error"
	if errors.Is(wrappedNotFound, ErrNotFound) {
		branch = "branch: user not found"
	}
	fmt.Fprintf(&out, "5) branch with errors.Is -> %s\n", branch) // branch: user not found

	// Тема: errors.As достаёт из цепочки ошибок конкретный тип.
	// Это полезно, когда ошибка хранит дополнительные данные.
	// В данном случае FieldError хранит Field и Reason.
	// Частая ошибка: сравнивать только текст ошибки вместо проверки типа через errors.As.
	fieldErr := &FieldError{
		Field:  "age",
		Reason: "must be positive",
	}
	wrappedField := fmt.Errorf("validation failed: %w", fieldErr)

	var target *FieldError
	ok := errors.As(wrappedField, &target)

	if ok {
		fmt.Fprintf(&out, "6) errors.As -> %t, field=%s, reason=%s\n", ok, target.Field, target.Reason) // true, age, must be positive
	} else {
		fmt.Fprintf(&out, "6) errors.As -> %t\n", ok)
	}

	// Специально неправильный пример:
	// Так делать НЕ нужно.
	// Нельзя завязывать бизнес-логику на точный текст ошибки.
	// Сегодня текст один, завтра туда добавят дополнительный контекст,
	// и сравнение строки сломается, хотя причина ошибки останется той же.
	textCompare := wrappedField.Error() == "age: must be positive"
	fmt.Fprintf(&out, "7) anti-example compare text -> %t\n", textCompare) // false

	return strings.TrimRight(out.String(), "\n")
}
