package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input  string
		output bool
	}{
		{"aa", true},
		{"ábá", true},
		{"Aba", true},
		{"ab", false},
		{"Abc", false},
	}
	for _, tc := range testCases {
		if IsPalindrome(tc.input) != tc.output {
			t.Errorf("IsPalindrome(%s) != %t", tc.input, tc.output)
		}
	}
}
