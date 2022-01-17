# go-enigma

```code
Usage: main.go ACTION_TYPE ALPHABET_TYPE CIPHER_NUM_1 CIPHER_NUM_2 CIPHER_NUM_3 CIPHER_NUM_CHANGE MESSAGE

ACTION_TYPE         - 'encode' or 'decode'
ALPHABET_TYPE       - Integer that will represent the alphaber type used. Values: 0 - ALPHA , 1 - NUMERIC, 2 - ALPHA_NUMERIC, 3 - ALPHA_NUMERIC_SYMBOL
CIPHER_NUM_1        - Chiper number 1
CIPHER_NUM_2        - Chiper number 2
CIPHER_NUM_3        - Chiper number 3
CIPHER_NUM_CHANGE   - A number used as a counter. When reached the cipher is changed.
MESSAGE             - Message to be ecoded / decoded

This can be used as a library now, by importing `github.com/heimbogdan/go-enigma/cipher`
```

## Usage

```go
//Used to set the alphabet that will be used, the cipher (three integers) and the counter used to update and change the cipher.
// AlphabetType: cipher.ALPHA (0), cipher.NUMERIC (1), cipher.ALPHA_NUMERIC (2), cipher.ALPHA_NUMERIC_SYMBOL (3)
cipher.SetScramble(aType AlphabetType, first int, second int, third int, counter int)

// Used to encode the message
cipher.Encode(message string)

// Used to decode the message
cipher.Decode(message string)
```
