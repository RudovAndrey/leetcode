package Golang

func longestCommonPrefix(strs []string) string {
	/*
		Обьявим массив байт, куда складываем префикс
		(Можно оптимизировать, создав срез байт с емкостью
		первой строки, таким образом избежим пересоздания базового массива)
		Берем первую строку за основу
		Идем по ней циклом
		Вложенным циклом пробегаемся по оставшимся строкам
		Если символ по индексу не равен, выходим
		Еслли длина строки меньше, чем текущий индекс, выходим
	*/
	var result []byte
	firstString := strs[0]
	for index := range firstString {
		for i := 1; i < len(strs); i++ {
			if len(strs[i]) < index+1 {
				return string(result)
			}
			if strs[i][index] != firstString[index] {
				return string(result)
			}
		}
		result = append(result, firstString[index])
	}
	return string(result)
}
