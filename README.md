# go-enigma

No main / executable function yet!

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
