package easy

func TwoSum(nums []int, target int) []int {
	for i, num1 := range nums {
		for j, num2 := range nums {
			if num1+num2 == target && i != j {
				return []int{i, j}
			}
		}
	}

	return []int{0, 0}
}
