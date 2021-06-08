package word

// IsPalindrome 判断字符串是否为回文
// 是，返回true，否，返回false
func IsPalindrome(letter string) bool {
	for i := range letter {
		if letter[i] != letter[len(letter)-i-1] {
			return false
		}
	}

	return true
}
