package palindrome

import (
	"testing"
)

func BenchmarkPalindrome(b *testing.B) {
	// If there is any complex initialization, perform it here and call b.ResetTimer to reset timer.
	// Therefore, it will not add to the measured time of each iteration.
	// time.Sleep(10 * time.Second)
	//b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsPalindrome("aaaa")
	}
}
