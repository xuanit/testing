package palindrome

import (
	"fmt"
)


func ExampleIsPalindrome()  {
	fmt.Printf("%t", IsPalindrome("aaaa"))
	// go test will check output is printed into standard output
	// Output:
	// true
}
