package easy

import "testing"

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		input  int
		expect bool
	}{
		{input: 1000021, expect: false},
		{input: 12134, expect: false},
		{input: 121, expect: true},
		{input: -121, expect: false},
		{input: 10, expect: false},
		{input: 1, expect: true},
	}

	for _, c := range cases {
		result := IsPalindrome(c.input)

		if result != c.expect {
			t.Fatalf("IsPalindrome(%v) = %v, expect: %v", c.input, result, c.expect)
		}
	}

}
