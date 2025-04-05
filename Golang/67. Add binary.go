package Golang

import (
	"strconv"
)

func addBinary(a string, b string) string {
	/*
		Вводим переменную аккумулятор, по классике CS
		Вводим результирующую строку
		Итак, решаем так. будем читать числа с конца
		Для этого определяем длины - 1, получаем конечные индексы
		Далее пробегаемся в цикле, пока один из индексов не равен нулю
		И прибавляем по следующим правилам:
		Если сумма цифр числа а, б и аккумулятора больше 1, значит
		в результрующую строку вводим эту сумму / 2
		В аккумулятор пойдет остаток от деления на 2.
		если сумма не больше 1, значит в результат вводим сумму,
		аккумулятор обнуляем.

		После этого мы определяем хвост (индекс хвоста у нас уже есть)

		Дальше мы итерируемся по этому хвосту, пока аккумулятор > 0
		если сумма аккумулятор + цифра > 1 значит в результат пишем 0
		иначе в результь пишем сумму, обноуляем аккумулятор.

		Если индекс > 0, прибавляем оставшийся хвост к началу строки.
		Выводим.
	*/
	accumulator := 0
	indexA, indexB := len(a)-1, len(b)-1
	var intA, intB, tailInt, sum, indexTail int
	var tail, result string
	for indexA >= 0 && indexB >= 0 {
		intA, _ = strconv.Atoi(string(a[indexA]))
		intB, _ = strconv.Atoi(string(b[indexB]))
		sum = intA + intB + accumulator
		if sum > 1 {
			result = strconv.Itoa(sum%2) + result
			accumulator = 1
		} else {
			result = strconv.Itoa(sum) + result
			accumulator = 0
		}
		indexA--
		indexB--
	}
	if len(a) > len(b) {
		tail = a
		indexTail = indexA
	} else {
		tail = b
		indexTail = indexB
	}
	for accumulator > 0 && indexTail >= 0 {
		tailInt, _ = strconv.Atoi(string(tail[indexTail]))
		sum = tailInt + accumulator
		if sum > 1 {
			result = strconv.Itoa(0) + result
		} else {
			result = strconv.Itoa(sum) + result
			accumulator = 0
		}
		indexTail--
	}

	if accumulator > 0 {
		result = "1" + result
	}
	if indexTail >= 0 {
		result = tail[:indexTail+1] + result
	}
	return result
}
