# classroom3: циклы, функции, массивы, слайсы

Классная работа для третьего занятия курса Go.

## Темы

1. `internal/loops` — `for`, `range`, `break`, `continue`, вложенные циклы.
2. `internal/functions` — параметры, возвращаемые значения, несколько return, variadic, function as value.
3. `internal/arrays` — массивы, фиксированная длина, value semantics, сравнение массивов.
4. `internal/slices` — `append`, `len`, `cap`, под-слайсы, копирование, удаление, вставка.

## Как запускать

```bash
make run-all
make test
make ci
```

## Главное правило структуры

`cmd/*/main.go` только вызывает `Example()`.
Вся логика находится в `internal/*`.
