package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder

	array := []rune(input)
	if len(array) == 0 {
		return "", nil
	}
	if unicode.IsDigit(array[0]) {
		return "", ErrInvalidString
	}
	for i, letter := range array {
		if !unicode.IsDigit(letter) {
			b.WriteRune(letter)
		} else if unicode.IsDigit(letter) {
			if unicode.IsDigit(array[i-1]) {
				return "", ErrInvalidString
			}
			k := int(letter - '0')
			if letter-'0' == 0 {
				k := b.String()
				b.Reset()
				b.WriteString(k[:len(k)-1])
			} else {
				b.WriteString(strings.Repeat(string(array[i-1]), k-1))
			}
		}
	}
	return b.String(), nil
}
