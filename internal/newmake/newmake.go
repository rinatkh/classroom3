package newmake

import "fmt"

type User struct {
	Name string
}

func Example() string {
	// Тема: new(T) выделяет память под zero value T и возвращает *T.
	// Частая ошибка: думать, что new возвращает само значение.
	age := new(int)
	*age = 26

	user := new(User)
	user.Name = "Maria"

	// Тема: make работает только для slice, map и chan.
	// make возвращает готовое значение, а не указатель.
	numbers := make([]int, 0, 3)  // [0 0 0]  len = 0 cap = 3
	numbers = append(numbers, 10) // [10 0 0] len = 1 cap = 3

	// Тема: nil slice и empty slice похожи по len, но отличаются по nil-проверке.
	var nilSlice []int           // память не выделена
	emptySlice := make([]int, 0) // память выделена

	// Ответ: new вернул указатели, make вернул рабочий слайс.
	return fmt.Sprintf("age=%d user=%s numbers=%v len=%d cap=%d nil=%t emptyNil=%t",
		*age, user.Name, numbers, len(numbers), cap(numbers), nilSlice == nil, emptySlice == nil)
}
