package Golang

/*
Алгоритм получения палиндрома такой
Вводим переменную digit, присваиваем ей последюю цифра
То, что осталось, присваиваем n1
Потом бежим в цикле, пока n1 >0
в каждой итерации берем последнюю цифру через остаток от деления на 10
Умножаем n1 на 10, увеличивая разрядность
Прибавляем digit к n1
Получаем число наоборот
*/
func reverseDigit(n1 int) int {
	digit := n1 % 10
	n1 /= 10
	n2 := digit
	for n1 > 0 {
		digit = n1 % 10
		n1 = n1 / 10
		n2 = n2 * 10
		n2 += digit
	}
	return n2
}

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	return x == reverseDigit(x)
}
