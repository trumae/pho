package pho

import "unicode"

func IsLetter(c interface{}) bool {
	r, ok := c.(rune)
	if !ok {
		return false
	}

	return unicode.IsLetter(r)
}
