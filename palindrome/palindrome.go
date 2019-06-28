package palindrome

import (
"unicode"
)

func IsPalindrome(s string) bool {
	var runes []rune
	for _ , r := range s {
		runes = append(runes, unicode.ToLower(r))
	}
	for i := range runes {
		if runes[i] != runes[len(runes) - 1 -i] {
			return false
		}
	}
	return true
}
