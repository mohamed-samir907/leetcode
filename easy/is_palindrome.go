package easy

import "fmt"

func IsPalindrome(x int) bool {
	nums := []byte(fmt.Sprint(x))
	lenght := len(nums)

	for i := 0; i <= lenght/2; i++ {
		if nums[i] != nums[lenght-(i+1)] {
			return false
		}
	}

	return true
}

func IsPalindromeFirstSolution(x int) bool {
	nums := []byte(fmt.Sprint(x))
	lenght := len(nums)
	center := int(lenght / 2)

	result := make(map[int]bool, center)

	// 12134
	// 0: 1 --> 5-(0+1) ---> 4: 4
	// 1: 2 --> 5-(1+1) ---> 3: 3
	// 2: 1 --> 5-(2+1) ---> 2: 1
	for i := 0; i <= center; i++ {
		start := nums[i]
		end := nums[lenght-(i+1)]

		if start == end {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	// check if the result map has a false value
	for _, value := range result {
		if !value {
			return false
		}
	}

	return true
}
