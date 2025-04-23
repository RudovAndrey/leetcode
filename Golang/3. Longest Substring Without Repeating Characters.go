package Golang

func lengthOfLongestSubstring_1(s string) int {
	/*
	   В голову приходит только наивное решение:
	   Двигаться с 0 символа, искать подстроки без повторящихся
	   Записывать длину найденной в max переменную
	   сложность - O(n^2)
	*/
	var max int
	for index, _ := range s {
		hashMap := map[byte]int{}
		i := index
		for i < len(s) {
			if _, exist := hashMap[s[i]]; exist {
				break
			} else {
				hashMap[s[i]] = 1
			}
			i++
		}
		result := i - index
		if result > max {
			max = result
		}
	}
	return max
}

func lengthOfLongestSubstring_2(s string) int {
	/*
				Решение, которое требуется - метод скользязего окна.
				Переменная max, которая хранит максимальную подстроку.
				hashMap, которая хранит уникальные символы
				Два указателя - правый и левый.
				Дивгаемся правым указателем, пока символы уникальны
				если сивол не уникален, двигаемся левым и удаляем сиволы по левому ндексу
		        до тех пор, пока неуникальный символ не уйдет из мапы. После того, как символы
				останутся только уникальными - определяем длину, если
				длина больше текущей, хранящейся в max - обновляем.
				"pwwkew" {p : true w: true}
	*/
	hashMap := map[byte]bool{}
	max_res, l, r := 0, 0, 0
	for r < len(s) {
		if _, exist := hashMap[s[r]]; exist {
			for {
				_, exist := hashMap[s[r]]
				if exist {
					delete(hashMap, s[l])
					l++
				} else {
					break
				}
			}
		}
		hashMap[s[r]] = true
		if r-l+1 > max_res {
			max_res = r - l + 1
		}
		r++
	}
	return max_res
}
