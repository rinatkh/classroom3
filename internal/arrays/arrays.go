package arrays

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// Тема: массив хранит фиксированное количество элементов одного типа.
	// [3]int и [4]int — разные типы, потому что длина входит в тип массива.
	numbers := [3]int{10, 20, 30}   // [10 20 30]
	first := numbers[0]             // 10
	last := numbers[len(numbers)-1] // 30

	fmt.Fprintf(&out, "first=%d\n", first)
	fmt.Fprintf(&out, "last=%d\n", last)
	fmt.Fprintf(&out, "len=%d\n", len(numbers)) // 3

	// Тема: zero value массива заполняет элементы zero value типа.
	// Для int это 0.
	var zero [2]int // [0 0]
	fmt.Fprintf(&out, "zero=%v\n", zero)

	// Тема: массив копируется при присваивании.
	// Частая ошибка: ожидать, что copyArray изменит original.
	original := [3]int{1, 2, 3} // [1 2 3]
	copyArray := original       // [1 2 3]
	copyArray[0] = 99           // [99 2 3]

	fmt.Fprintf(&out, "original=%v\n", original)
	fmt.Fprintf(&out, "copy=%v\n", copyArray)

	// Тема: массивы можно сравнивать, если элементы сравнимы.
	equal := [2]string{"go", "sql"} == [2]string{"go", "sql"} // true
	fmt.Fprintf(&out, "equal=%t", equal)

	return out.String()
}
