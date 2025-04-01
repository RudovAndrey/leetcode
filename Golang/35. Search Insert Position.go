package leetcode

func searchInsert(nums []int, target int) int {
	/*
		Решаем через бинарный поиск
		Если индекс равен = выдаем индекс
		Если значение по текущему индексу больше целевого и
		целевое значение больше, чем предыдущее от индекса, значит ставим элемент туда
		Сложность - O(log n)
	*/
	if target <= nums[0] {
		return 0
	}
	leftBorder, rightBorder := 1, len(nums)-1
	var index int
	for leftBorder <= rightBorder {
		index = (leftBorder + rightBorder) / 2
		if nums[index] == target || (target > nums[index-1] && target < nums[index]) {
			return index
		}
		if nums[index] > target {
			rightBorder = index - 1
		} else {
			leftBorder = index + 1
		}
	}
	return len(nums)
}
