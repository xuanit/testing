package palindrome

import (
	"os"
	"testing"
)

func TestSubpackagePalindrome(t *testing.T) {
	//t.Errorf("expected 1, but go %d", 3)
}

func TestSubpackagePalindrome2(t *testing.T) {
	//t.Errorf("expect 2, but got %d", 3)
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
