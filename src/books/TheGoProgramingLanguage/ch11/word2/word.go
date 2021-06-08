package word2

import (
	"unicode"
)


// IsPalindrome 判断字符串是否为回文
// 是，返回true，否，返回false
func IsPalindrome(letter string) bool {
	// 只支持可以打印的字符，并且转换成小写
	var letters []rune
	for _, val := range letter {
		if unicode.IsLetter(val) {
			letters = append(letters, unicode.ToLower(val))
		}
	}

	// 匹配处理后的字符
	for i := range letters {
		if letters[i] != letters[len(letters)-i-1] {
			return false
		}
	}

	return true
}
