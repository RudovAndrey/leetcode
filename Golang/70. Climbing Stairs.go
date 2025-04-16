package Golang

func climbStairs(n int, cache ...map[int]int) int {
	/*
	   Если ступенька одна, то способ один - 1
	   Если ступеньки две то, способа два - 1+1,2
	   если ступенки три, то способа три - 1+1+1,1+2,2+1
	   если ступеньки четыре, то способов уже пять - 1+1+1+1, 1+2+1, 2+1+1, 1+1+2
	   То есть n ступенек = n-1 способов + n-2 способов
	   Похоже на вычисление чисел фибоначчи. Чтобы избежать повторных вычислений,
	   нужно сделать мапку с количеством способов.
	*/
	if n <= 2 {
		return n
	}
	if len(cache) < 1 {
		cache = append(cache, map[int]int{0: 0})
	}
	if _, exist := cache[0][n]; exist {
		return cache[0][n]
	}
	cache[0][n] = climbStairs(n-1, cache[0]) + climbStairs(n-2, cache[0])
	return cache[0][n]
}
