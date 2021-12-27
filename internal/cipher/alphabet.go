package cipher

import (
	"fmt"
	"strings"
)

type baseAlphabet struct {
	encodeMap map[int]int
	decodeMap map[int]int
}

var cipher1 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}
var cipher2 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}
var cipher3 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}

var alpha = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var numeric = "0123456789"
var alphaNumeric = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
var alphaNumericSymbols = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ -_,.;:'\"[]{}()=+<>!@#$%^&*`~"

var _first = int(0)
var _second = int(0)
var _third = int(0)
var _counter = int(0)
var _aType = ALPHA
var _aSize = len(alpha)

type AlphabetType int8

const (
	ALPHA AlphabetType = iota
	NUMERIC
	ALPHA_NUMERIC
	ALPHA_NUMERIC_SYMBOLS
)

// f 10
func SetScramble(aType AlphabetType, first int, second int, third int, counter int) {
	_first = first
	_second = second
	_third = third
	_counter = counter
	_aType = aType
	switch aType {
	case ALPHA:
		_aSize = len(alpha)
	case NUMERIC:
		_aSize = len(numeric)
	case ALPHA_NUMERIC:
		_aSize = len(alphaNumeric)
	case ALPHA_NUMERIC_SYMBOLS:
		_aSize = len(alphaNumericSymbols)
	}
}

func scramble1(num int) {
	resetCipher1()
	for i := 1; i <= _aSize; i++ {
		index := (((num % i) + num + i - 1) % _aSize) + 1
		for cipher1.decodeMap[index] != 0 {
			index = (index % _aSize) + 1
		}
		cipher1.encodeMap[i] = index
		cipher1.decodeMap[index] = i
		num = index
	}
}

func scramble2(num int) {
	resetCipher2()
	for i := 1; i <= _aSize; i++ {
		index := (((num % i) + num + i - 1) % _aSize) + 1
		for cipher2.decodeMap[index] != 0 {
			index = (index % _aSize) + 1
		}
		cipher2.encodeMap[i] = index
		cipher2.decodeMap[index] = i
		num = index
	}
}

func scramble3(num int) {
	resetCipher3()
	for i := 1; i <= _aSize; i++ {
		index := (((num % i) + num + i - 1) % _aSize) + 1
		for cipher3.decodeMap[index] != 0 {
			index = (index % _aSize) + 1
		}
		cipher3.encodeMap[i] = index
		cipher3.decodeMap[index] = i
		num = index
	}
}

func resetCipher1() {
	cipher1 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}
}

func resetCipher2() {
	cipher2 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}
}

func resetCipher3() {
	cipher3 = baseAlphabet{encodeMap: make(map[int]int), decodeMap: make(map[int]int)}
}

func PrintCipher() {
	fmt.Println("Cipher1: ")
	for i := 1; i <= _aSize; i++ {
		fmt.Printf(" %d => %d |", i, cipher1.encodeMap[i])
	}
	fmt.Println("\nCipher2: ")
	for i := 1; i <= _aSize; i++ {
		fmt.Printf(" %d => %d |", i, cipher2.encodeMap[i])
	}
	fmt.Println("\nCipher3: ")
	for i := 1; i <= _aSize; i++ {
		fmt.Printf(" %d => %d |", i, cipher3.encodeMap[i])
	}
}

func getAlphabet(aType AlphabetType) string {
	switch aType {
	case ALPHA:
		return alpha
	case NUMERIC:
		return numeric
	case ALPHA_NUMERIC:
		return alphaNumeric
	case ALPHA_NUMERIC_SYMBOLS:
		return alphaNumericSymbols
	default:
		return ""
	}
}

func Encode(message string) string {
	scramble1(_first)
	scramble2(_second)
	scramble3(_third)
	encMessage := make([]rune, len(message))
	alphabet := getAlphabet(_aType)
	alphabetRunes := []rune(alphabet)
	counter := int(0)
	c1 := int(0)
	c2 := int(0)
	c3 := int(0)
	for i, ch := range message {
		index := strings.IndexRune(alphabet, ch)
		if index < 0 {
			encMessage[i] = ch
		} else {
			i1 := cipher1.encodeMap[index+1]
			i2 := cipher2.encodeMap[i1]
			i3 := cipher3.encodeMap[i2]
			encMessage[i] = alphabetRunes[i3-1]
		}
		counter++
		if counter == _counter {
			counter = int(0)
			c1++
			scramble1(((_first + c1) % _aSize) + 1)
			if c1 == _aSize {
				c1 = int(0)
				c2++
				scramble2(((_second + c2) % _aSize) + 1)
				if c2 == _aSize {
					c2 = int(0)
					c3++
					scramble3(((_third + c3) % _aSize) + 1)
					if c3 == _aSize {
						c3 = int(0)
					}
				}
			}
		}
	}
	return string(encMessage)
}

func Decode(message string) string {
	scramble1(_first)
	scramble2(_second)
	scramble3(_third)
	decMessage := make([]rune, len(message))
	alphabet := getAlphabet(_aType)
	alphabetRunes := []rune(alphabet)
	counter := int(0)
	c1 := int(0)
	c2 := int(0)
	c3 := int(0)
	for i, ch := range message {
		index := strings.IndexRune(alphabet, ch)
		if index < 0 {
			decMessage[i] = ch
		} else {
			i3 := cipher3.decodeMap[index+1]
			i2 := cipher2.decodeMap[i3]
			i1 := cipher1.decodeMap[i2]
			decMessage[i] = alphabetRunes[i1-1]
		}
		counter++
		if counter == _counter {
			counter = int(0)
			c1++
			scramble1(((_first + c1) % _aSize) + 1)
			if c1 == _aSize {
				c1 = int(0)
				c2++
				scramble2(((_second + c2) % _aSize) + 1)
				if c2 == _aSize {
					c2 = int(0)
					c3++
					scramble3(((_third + c3) % _aSize) + 1)
					if c3 == _aSize {
						c3 = int(0)
					}
				}
			}
		}
	}
	return string(decMessage)
}
