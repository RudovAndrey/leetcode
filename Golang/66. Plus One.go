package Golang

func plusOne(digits []int) []int {
	/*
	   Задачка немного проще, чем такая же, но с двоичными числами
	   Вводим переменную inMind, куда переносим единицу следующего разряда
	   Первым действием прибавляем к последнему единицу
	   если получилось 10, то в inMind записываем единицу
	   Если нет - то просто +1 к числу
	   далее пробегаемся по оставшемуся массиву до тех по пока он не кончится
	   или пока в inMind не останется ноль.
	   Далее, если в inMind все еще осталась единица, вставляем ее на первое место
	   Сложность - O(n) память - еще один массив в последнем действии, создается новый
	   базовый массив
	*/
	var inMind int
	lastNum := digits[len(digits)-1] + 1
	if lastNum < 10 {
		digits[len(digits)-1] = lastNum
	} else {
		digits[len(digits)-1] = 0
		inMind = 1
	}
	index := len(digits) - 2
	for inMind > 0 && index >= 0 {
		if digits[index]+inMind == 10 {
			digits[index] = 0
		} else {
			digits[index] = digits[index] + inMind
			inMind = 0
		}
		index--
	}
	if inMind == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
