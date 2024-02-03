package gomaith

// Slenderize changes the last vowel in s to be a slender vowel.
func Slenderize(s string) string {
	ss := []rune(s)
	ss = slenderizeSlice(ss)
	return string(ss)
}

func slenderizeSlice(ss []rune) []rune {
	if len(ss) < 2 {
		return ss
	}
	if _, found := irishConsonants[ss[len(ss)-1]]; !found {
		return ss
	}
	lastVowelIndex := lastIndex(ss, func(r rune) bool {
		_, found := irishVowels[r]
		return found
	})

	// Only add if the last broad vowel is not at the beginning or at the end of the word.
	switch lastVowelIndex {
	case -1, 0, len(ss) - 1:
		return ss
	}

	switch ss[lastVowelIndex] {
	case 'a', 'á', 'o', 'ó', 'u', 'ú':
		ss = append(ss[:lastVowelIndex+1], append([]rune{'i'}, ss[lastVowelIndex+1:]...)...)
	}
	return ss
}
