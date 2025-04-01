package twoSum

func twoSum(nums []int, target int) []int {
	/*
		Строим Hash map, проходя по каждому элементу
		Если в хэшмапе нет элемента, равному разности target - текущий элемент
		То мы заносим туда элемент: его индекс. Если есть, то выбрасываем индекс
		текущий + индекс найденный в мапе.
		Сложность - O(n), + затраты по памяти на мапу.
	*/
	hashMap := map[int]int{}
	for index, elem := range nums {
		if _, exist := hashMap[target-elem]; exist {
			return []int{hashMap[target-elem], index}
		} else {
			hashMap[elem] = index
		}
	}
	return []int{}
}
