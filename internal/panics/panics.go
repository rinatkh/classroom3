package panics

import (
	"fmt"
	"strings"
)

func Example() string {
	var out strings.Builder

	// Тема: recover, вызванный вне deferred-функции, ничего не поймает.
	// Частая ошибка: вызывать recover в обычном коде и ждать, что panic будет перехвачен.
	// Переменная outside = nil
	outside := recover() // nil
	fmt.Fprintf(&out, "1) recover outside defer -> %v\n", outside)

	// Тема: defer выполняется и при обычном выходе из функции.
	// Здесь panic нет, поэтому код доходит до конца, а потом выполняется defer.
	// Переменная noPanicFlow = "start -> end -> defer cleanup"
	noPanicFlow := func() (result string) {
		var log strings.Builder

		defer func() {
			fmt.Fprintf(&log, " -> defer cleanup")
			result = log.String()
		}()

		fmt.Fprintf(&log, "start -> end")
		return
	}()

	fmt.Fprintf(&out, "2) without panic -> %s\n", noPanicFlow)

	// Тема: panic — аварийное завершение текущего потока выполнения.
	// После panic обычный код ниже уже не выполнится.
	// Но defer всё равно выполнится.
	//
	// Тема: recover может поймать panic только внутри deferred-функции.
	// Здесь сначала сработает defer с recover, а потом defer cleanup,
	// потому что defer выполняются в обратном порядке — LIFO.
	//
	// Переменная panicFlow =
	// "start -> before panic -> recover=unexpected state -> defer cleanup"
	panicFlow := func() (result string) {
		var log strings.Builder

		defer func() {
			fmt.Fprintf(&log, " -> defer cleanup")
			result = log.String()
		}()

		defer func() {
			if r := recover(); r != nil {
				fmt.Fprintf(&log, " -> recover=%v", r)
			}
		}()

		fmt.Fprintf(&log, "start -> before panic")

		panic("unexpected state")

		// Так делать НЕ нужно:
		// код после panic в этой функции уже не выполнится.
		fmt.Fprintf(&log, " -> after panic")
		return
	}()

	fmt.Fprintf(&out, "3) with panic -> %s", panicFlow)

	return out.String()
}
