package Golang

func lengthOfLastWord(s string) int {
	/*
		Задача простая, но строка может содержать пробелы в конце
		пробелы можно убрать через TrimSpace, но это сходу увеличит
		сложность, попробуем справиться без этого метода.
		Введем флаг пробелов
		Введем переменную для количества пробелов в конце
		В цилке читаем строку с конца
		Первое условие - проверяем, что пробелы в конце закончились
		если закончились - ставим флаг в true, количество пробелов
		равняется длине минус текущий индекс и пропускаем остаток цикла
		Далее уже ищем когда будет следующий пробел после символов и выдаем
		длину - индекс - количество пробелов
		Если пробелов не было, просто вернем длину строки
	*/
	var trailingSpaceFlag bool
	var spaceCount int
	for i := len(s) - 1; i >= 0; i-- {
		if string(s[i]) != " " && trailingSpaceFlag == false {
			trailingSpaceFlag = true
			spaceCount = len(s) - i
			continue
		}
		if string(s[i]) == " " && trailingSpaceFlag == true {
			return len(s) - i - spaceCount
		}
	}
	return len(s)
}
