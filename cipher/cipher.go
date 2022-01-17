package cipher

import c "github.com/heimbogdan/go-enigma/internal/cipher"

const (
	ALPHA c.AlphabetType = iota
	NUMERIC
	ALPHA_NUMERIC
	ALPHA_NUMERIC_SYMBOLS
)

func SetScramble(aType c.AlphabetType, first int, second int, third int, counter int) {
	c.SetScramble(c.ALPHA_NUMERIC_SYMBOLS, 11, 25, 3, 3)
}

func Encode(message string) string {
	return c.Encode(message)
}

func Decode(message string) string {
	return c.Decode(message)
}
