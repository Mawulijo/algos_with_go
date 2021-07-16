package main

func SumNumsInList(list []int) int {
	var sum int
	for _, v := range list {
		sum += v
	}
	return sum
}

// SumNumsInListR uses a recursive approach
func SumNumsInListR(list []int) int {
	if len(list) == 0 {
		return 0
	}
	return list[0] + SumNumsInListR(list[1:])
}