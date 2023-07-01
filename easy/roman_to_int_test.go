package easy

import "testing"

func TestIsRomanToInt(t *testing.T) {
	cases := []struct {
		input  string
		expect int
	}{
		{input: "III", expect: 3},
		{input: "LVIII", expect: 58},
		{input: "MCMXCIV", expect: 1994},
		{input: "IV", expect: 4},
		{input: "IVLV", expect: 59},
	}

	for _, c := range cases {
		result := RomanToInt(c.input)

		if result != c.expect {
			t.Fatalf("RomanToInt(%v) = %v, expect: %v", c.input, result, c.expect)
		}
	}

}
