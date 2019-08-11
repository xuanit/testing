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
      t.Parallel()
			if IsPalindrome(tc.input) != tc.output {
				t.Errorf("IsPalindrome(%s) != %t", tc.input, tc.output)
			}
		})

	}
}
