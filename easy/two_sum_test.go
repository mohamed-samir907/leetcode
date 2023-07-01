package easy

import "testing"

func TestTwoSum(t *testing.T) {
	cases := []struct {
		nums   []int
		target int
		expect []int
	}{
		{
			nums:   []int{2, 7, 11, 15},
			target: 9,
			expect: []int{0, 1},
		},
		{
			nums:   []int{3, 2, 4},
			target: 6,
			expect: []int{1, 2},
		},
		{
			nums:   []int{3, 3},
			target: 6,
			expect: []int{0, 1},
		},
	}

	for _, c := range cases {
		result := TwoSum(c.nums, c.target)

		if result[0] != c.expect[0] && result[1] != c.expect[1] {
			t.Fatalf("TwoSum(%v, %v) = (%v, %v), expect (%v, %v)",
				c.nums, c.target,
				result[0], result[1],
				c.expect[0], c.expect[1])
		}
	}

}
