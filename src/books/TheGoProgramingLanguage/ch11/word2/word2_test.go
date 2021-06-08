package word2

import (
	"testing"
)

func TestWord(t *testing.T) {
	// 表格驱动测试
	testSuits := []struct {
		Input    string
		Expected bool
	}{
		{"abcba", true},
		{"abbacbba", false},
		{"A man, a plan, a canal: Panama", true},
		{"Aba", true},
	}

	for _, suit := range testSuits {
		if actual := IsPalindrome(suit.Input); actual != suit.Expected {
			// f(x) = y，want z(如果有必要)这种格式的表达格式很常见
			t.Errorf("IsPalindrome(%s) = %t, Expected:%t", suit.Input, actual, suit.Expected)
		}
	}

}
