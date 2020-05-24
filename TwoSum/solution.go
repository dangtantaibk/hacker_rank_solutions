func twoSum(nums []int, target int) []int {
	mapReturn := make(map[int]int)
	for k, v := range nums {
		numNeed := target - v
		if value, ok := mapReturn[numNeed]; ok {
			return []int{value, k}
		} else {
			mapReturn[v] = k
		}
	}
	return []int{}
}