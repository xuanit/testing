package palindrome

import (
	"testing"
)

func BenchmarkPalindrome(b *testing.B)  {
	// If there is any complex initialization, perform it here and call b.ResetTimer to reset timer.
	// Therefore, it will not add to the measured time of each iteration.
	for i := 0; i < b.N; i++ {
		IsPalindrome("aaaa")
	}
}

func TestPalindrome(t *testing.T)  {
	testCases := []struct{
		input string
		output bool
	}{
	{"aa", true},
		{"ábá", true},
		{"Aba", true},
		{"ab", false},
		{ "Abc", false},
	}
	for _, tc := range testCases {
		if IsPalindrome(tc.input) != tc.output {
			t.Errorf("IsPalindrome(%s) != %t", tc.input, tc.output)
		}
	}
}
