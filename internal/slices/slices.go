package slices

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// Тема: слайс — это descriptor: pointer на массив, len и cap.
	base := []int{10, 20, 30, 40, 50} // [10 20 30 40 50], len=5, cap=5
	part := base[1:4]                 // [20 30 40], len=3, cap=4

	// Тема: part смотрит на тот же underlying array.
	// Частая ошибка: думать, что part — независимая копия.
	part[0] = 99 // base=[10 99 30 40 50], part=[99 30 40]

	// Тема: append может переиспользовать массив или создать новый.
	// Результат append всегда нужно присваивать.
	part = append(part, 77) // base=[10 99 30 40 77], part=[99 30 40 77], len=4, cap=4

	// Тема: copy создаёт независимую копию значений.
	clone := make([]int, len(part)) // [0 0 0 0]
	copy(clone, part)               // [99 30 40 77]
	clone[0] = 555                  // [555 30 40 77]

	fmt.Fprintf(&out, "base=%v\n", base)
	fmt.Fprintf(&out, "part=%v\n", part)
	fmt.Fprintf(&out, "clone=%v\n", clone)
	fmt.Fprintf(&out, "len=%d\n", len(part))
	fmt.Fprintf(&out, "cap=%d", cap(part))

	return out.String()
}

func AppendMemoryExample() string {
	var out strings.Builder

	// Тема: append добавляет элементы в слайс.
	// len показывает, сколько элементов видно через слайс.
	// cap показывает, сколько элементов можно использовать до выделения нового массива.
	s := make([]int, 0, 3) // [], len=0, cap=3
	fmt.Fprintf(&out, "after make -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

	// Тема: пока capacity хватает — append использует тот же базовый массив.
	s = append(s, 10) // [10], len=1, cap=3
	fmt.Fprintf(&out, "after append 10 -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

	s = append(s, 20) // [10 20], len=2, cap=3
	fmt.Fprintf(&out, "after append 20 -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

	s = append(s, 30) // [10 20 30], len=3, cap=3
	fmt.Fprintf(&out, "after append 30 -> s=%v len=%d cap=%d\n", s, len(s), cap(s))

	// Тема: когда capacity не хватает — append создаёт новый массив,
	// копирует старые элементы и возвращает новый слайс.
	s = append(s, 40) // [10 20 30 40], len=4, cap=6
	fmt.Fprintf(&out, "after append 40 -> s=%v len=%d cap=%d", s, len(s), cap(s))

	return out.String()
}

func change(part []int) string {
	var out strings.Builder

	// Тема: в функцию передаётся не весь массив, а slice header.
	// Он содержит pointer, len и cap, но смотрит на тот же базовый массив.
	before := fmt.Sprintf("part=%v len=%d cap=%d", part, len(part), cap(part)) // [20 30], len=2, cap=3
	fmt.Fprintf(&out, "%s\n", before)

	// Тема: изменение part[0] меняет общий базовый массив.
	part[0] = 999                                                                           // numbers=[10 999 30 40], part=[999 30]
	afterSet := fmt.Sprintf("after set: part=%v len=%d cap=%d", part, len(part), cap(part)) // [999 30], len=2, cap=3
	fmt.Fprintf(&out, "%s\n", afterSet)

	// Тема: первый append ещё помещается в capacity,
	// поэтому запись идёт в старый базовый массив.
	part = append(part, 777)                                                                        // part=[999 30 777], len=3, cap=3
	afterAppend1 := fmt.Sprintf("after append1: part=%v len=%d cap=%d", part, len(part), cap(part)) // [999 30 777], len=3, cap=3
	fmt.Fprintf(&out, "%s\n", afterAppend1)

	// Тема: второй append уже не помещается в capacity,
	// поэтому создаётся новый массив.
	part = append(part, 888)                                                                        // part=[999 30 777 888], len=4, cap=6
	afterAppend2 := fmt.Sprintf("after append2: part=%v len=%d cap=%d", part, len(part), cap(part)) // [999 30 777 888], len=4, cap=6
	fmt.Fprintf(&out, "%s", afterAppend2)

	return out.String()
}

func SliceInFunctionExample() string {
	var out strings.Builder

	// Тема: срез numbers[1:3] создаёт новый slice header.
	// len = high - low = 3 - 1 = 2
	// cap = cap(numbers) - low = 4 - 1 = 3
	numbers := []int{10, 20, 30, 40} // [10 20 30 40], len=4, cap=4

	log := change(numbers[1:3])

	// Ответ: первый append меняет исходный numbers,
	// а второй append создаёт новый массив и уже не меняет numbers.
	fmt.Fprintf(&out, "%s\n", log)
	fmt.Fprintf(&out, "numbers=%v", numbers)

	return out.String()
}
