package Golang

func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
	   В общем, пришлось подсмотреть решение.
	   Все будем решать с конца. Определяем последний элемент в первом списке,
	   который будет результирующим, опрделеляем длины микромассивов
	   И начинаем собирать массив с конца. Жаль, что сам не допер.
	   Пробегаемся, пока есть элементы в массивах
	   После это в цикле пробегаемся по оставшемуся nums2 и просто дописываем его
	   в начало по индексу mainIndex
	*/

	mainIndex := n + m - 1
	index1, index2 := m-1, n-1
	for index1 >= 0 && index2 >= 0 {
		if nums1[index1] > nums2[index2] {
			nums1[mainIndex] = nums1[index1]
			index1--
		} else {
			nums1[mainIndex] = nums2[index2]
			index2--
		}
		mainIndex--
	}
	for index2 >= 0 {
		nums1[mainIndex] = nums2[index2]
		index2--
		mainIndex--
	}
}
