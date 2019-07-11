package palindrome

import (
	"fmt"
	"testing"
)

func BenchmarkPalindrome(b *testing.B)  {
	// If there is any complex initialization, perform it here and call b.ResetTimer to reset timer.
	// Therefore, it will not add to the measured time of each iteration.
	for i := 0; i < b.N; i++ {
		IsPalindrome("aaaa")
	}
}

func TestIsPalindrome(t *testing.T)  {
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

func ExampleIsPalindrome()  {
	fmt.Printf("%t", IsPalindrome("aaaa"))
	// go test will check output is printed into standard output
	// Output:
	// true
}
