package gomaith

import (
	"strings"
	"unicode"
)

var irishLetters = M('a', 'á', 'b', 'c', 'd', 'e', 'é', 'f', 'g', 'h', 'i', 'í', 'l', 'm', 'n', 'o', 'ó', 'p', 'r', 's', 't', 'u', 'ú')
var irishVowels = M('a', 'á', 'e', 'é', 'i', 'í', 'o', 'ó', 'u', 'ú')
var irishConsonants = difference(irishVowels, irishLetters)
var lenitable = M('b', 'c', 'd', 'f', 'g', 'm', 'p', 's', 't')

// Vocative takes a string and returns the vocative form of it.
// The vocative form is used when addressing someone directly.
// For example, "Seán" becomes "a Sheáin" in the vocative.
// The rules for forming the vocative are as follows:
// 0. Add the vocative particle "a " to the beginning of the name.
// 1. If this is not an Irish name, just return as-is.
// 2. Names that start with a consonant get a séimhiú 'h'.
func Vocative(s string) string {
	// Empty strings get returned as-is.
	if s == "" {
		return s
	}

	// Change to a []rune for easier manipulation.
	ss := []rune(strings.ToLower(s))

	// 1. If this is not an Irish name, just return as-is.
	// Determining whether a name is Irish is a hard problem, but for the
	// purposes of this exercise we will assume that any string containing
	// characters not found in the Irish alphabet is not an Irish name.
	for _, r := range ss {
		if _, found := irishLetters[r]; !found {
			// Capitalize the first letter of the string after we lower-cased it.
			ss[0] = unicode.ToUpper(ss[0])
			// Finally, add the vocative particle "a " to the beginning of the string.
			ss = append([]rune("a "), ss...)
			return string(ss)
		}
	}

	// 2. Names that start with a consonant get a séimhiú 'h' appended to the first rune.
	// Lenitable consonants in Irish are b, c, d, f, g, m, p, s, and t.
	ss = leniteSlice(ss)

	// 3. If the name ends in a broad consonant, slenderise it by adding
	// an 'i' just before it. A broad consonant is one that comes after a, o, or u.
	// Exception: this is only done for proper nouns, and not common nouns.
	// Named entity recognition is a hard problem, so for the purposes of this
	// exercise we will assume that all capitalized strings are proper nouns.
	if _, found := irishConsonants[ss[len(ss)-1]]; found {
		lastVowelIndex := lastIndex(ss, func(r rune) bool {
			_, found := irishVowels[r]
			return found
		})
		// Only add if the last broad vowel is not at the beginning or at the end of the word.
		if lastVowelIndex != -1 && lastVowelIndex != 0 && lastVowelIndex != len(ss)-1 {
			switch ss[lastVowelIndex] {
			case 'a', 'á', 'o', 'ó', 'u', 'ú':
				ss = append(ss[:lastVowelIndex+1], append([]rune{'i'}, ss[lastVowelIndex+1:]...)...)
			default: // do nothing if it is not a broad vowel
			}
		}
	}

	// Capitalize the first letter of the string after we lower-cased it.
	ss[0] = unicode.ToUpper(ss[0])
	// Finally, add the vocative particle "a " to the beginning of the string.
	ss = append([]rune("a "), ss...)
	return string(ss)
}
