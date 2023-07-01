package easy

func RomanToInt(s string) int {
	nums := map[rune]int{
		'M': 1000,
		'D': 500,
		'C': 100,
		'L': 50,
		'X': 10,
		'V': 5,
		'I': 1,
	}

	length := len(s)

	sum := 0
	for i := 0; i < length; i++ {
		value := rune(s[i])

		if i == length-1 {
			sum += nums[value]
			continue
		}

		next := rune(s[i+1])

		if nums[next] > nums[value] {
			sum += (nums[next] - nums[value])
			i++ // skip the next value
		} else {
			sum += nums[value]
		}
	}

	return sum
}
