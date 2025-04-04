package Golang

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
