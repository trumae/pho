package pho

import "unicode"

func IsLetter(c interface{}) bool {
	r, ok := c.(rune)
	if !ok {
		return false
	}

	return unicode.IsLetter(r)
}

func IsDigit(c interface{}) bool {
	r, ok := c.(rune)
	if !ok {
		return false
	}

	return unicode.IsDigit(r)
}

func IsSpace(c interface{}) bool {
	r, ok := c.(rune)
	if !ok {
		return false
	}

	return unicode.IsSpace(r)
}
