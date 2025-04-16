package Golang

func mySqrt(x int) int {
	/*
	   Возьмем за основу метод половинного деления.
	   Точность определена, число целое
	   например, возьмем 6
	   Пределы - 0 и 6
	   вычисляем медиану 0+6/2, получаем 3, 3*3 больше 6, предел 0 3
	   Следующая итерация вычисляем медиану 0+3/2 получаем 1 1*1 меньше 6, предел 1 3
	   Следующая - вычисляем медиану 1+3/2 получаем 2, 2*2 меньше 6, предел 2 3
	   Следущая - вычисляем медиану 2+3/2

	*/
	if x < 2 {
		return x
	}
	min := 0
	max := x
	delimiter := x / 2
	for delimiter*delimiter != x {
		delimiter = (min + max) / 2
		if delimiter == min {
			return delimiter
		}
		if delimiter*delimiter > x {
			max = delimiter
		} else {
			min = delimiter
		}
	}
	return delimiter
}
