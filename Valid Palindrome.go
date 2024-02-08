func isPalindrome(s string) bool {
	sb := []byte(s)
	l, r := 0, len(s)-1

	for l < r {
		if isAlpha(s[l]) == false {
			l++
			continue
		}
		if isAlpha(s[r]) == false {
			r--
			continue
		}
		if strings.ToLower(string(sb[l])) != strings.ToLower(string(sb[r])) {
			return false
		}
		l++
		r--
	}
	return true
}

func isAlpha(x byte) bool {
	if (x >= 65 && x <= 90) || (x >= 97 && x <= 122) {
		return true
	}
	return false
}