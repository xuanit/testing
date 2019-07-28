package palindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		output bool
	}{
		{"ShouldHandleAsciiPalindrome", "aa", true},
		{"ShouldHandleAsciiNonPalindrome", "ab", false},
		{"ShouldHandleUnicodePalindrome", "ábá", true},
		{"ShouldHandlePalindromeWithUpperCase", "Aba", true},
		{"shouldHandleNonPalindromeWithUpperCase", "Abc", false},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if IsPalindrome(tc.input) != tc.output {
				t.Parallel()
				t.Errorf("IsPalindrome(%s) != %t", tc.input, tc.output)
			}
		})

	}
}
