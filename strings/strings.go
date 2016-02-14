package strings

// RemoveIf removes all elements satisfying specific criteria and returns a new string
func RemoveIf(s string, f func(rune) bool) string {
	runes := []rune(s)
	result := 0
	for i, r := range runes  {
		if !f(r) {
			runes[result] = runes[i]
			result++
		}
	}

	return string(runes[0:result])
}
