package palindrome

import (
	"fmt"
)

func ExampleIsPalindrome() {
	fmt.Printf("%t\n", IsPalindrome("aaaa"))
	fmt.Printf("%t\n", IsPalindrome("aaac"))
	// Output:
	// true
	// false
}
