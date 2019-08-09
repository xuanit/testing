package palindrome

import (
	"fmt"
	"testing"
)

func BenchmarkSubpackageParallel(b *testing.B) {
	// If there is any complex initialization, perform it here and call b.ResetTimer to reset timer.
	// Therefore, it will not add to the measured time of each iteration.
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("Hello %s!", "world")

		}
	})
}
