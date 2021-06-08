package word

import "testing"

func TestWord(t *testing.T) {
	word := "abcdedcba"
	if !IsPalindrome(word) {
		t.Errorf("%s is Palindrome, expected %t", word, true)
	}

	word2 := "eoe"
	if !IsPalindrome(word2) {
		t.Errorf("%s is Palindrome, expected %t", word2, true)
	}
}

func TestNoWord(t *testing.T) {
	word := "adsfjalskdvj"
	if IsPalindrome(word) {
		t.Errorf("%s is Palindrome, expected %t", word, true)
	}

	word = "adsfjalskdvj"
	if IsPalindrome(word) {
		t.Errorf("%s is not Palindrome, expected %t", word, true)
	}
}
