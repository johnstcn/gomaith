package gomaith

// Lenite adds a séimhiú to the beginning of s if it begins with a lenitable
// consonant and does not already have a séimhiú after the initial consonant.
func Lenite(s string) string {
	ss := []rune(s)
	return string(leniteSlice(ss))
}

func leniteSlice(ss []rune) []rune {
	if len(ss) < 2 {
		return ss
	}
	if _, found := lenitable[ss[0]]; !found {
		return ss
	}
	if ss[1] == 'h' {
		return ss
	}
	ss = append([]rune{ss[0], 'h'}, ss[1:]...)
	return ss
}
