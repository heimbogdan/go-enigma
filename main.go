package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/heimbogdan/go-enigma/cipher"
	c "github.com/heimbogdan/go-enigma/internal/cipher"
)

func main() {
	if len(os.Args) < 8 {
		printHelp()
	} else {
		action := os.Args[1]
		aType, _ := strconv.Atoi(os.Args[2])
		c1, _ := strconv.Atoi(os.Args[3])
		c2, _ := strconv.Atoi(os.Args[4])
		c3, _ := strconv.Atoi(os.Args[5])
		cc, _ := strconv.Atoi(os.Args[6])
		cipher.SetScramble(c.AlphabetType(aType), c1, c2, c3, cc)
		if action == "encode" {
			fmt.Println(cipher.Encode(os.Args[7]))
		} else if action == "decode" {
			fmt.Println(cipher.Decode(os.Args[7]))
		} else {
			fmt.Println("Wrong ACTION_TYPE!")
			printHelp()
		}
	}
}

func printHelp() {
	fmt.Println("Usage: main.go ACTION_TYPE ALPHABET_TYPE CIPHER_NUM_1 CIPHER_NUM_2 CIPHER_NUM_3 CIPHER_NUM_CHANGE MESSAGE")
	fmt.Println("\nACTION_TYPE         - 'encode' or 'decode'")
	fmt.Println("ALPHABET_TYPE       - Integer that will represent the alphaber type used. Values: 0 - ALPHA , 1 - NUMERIC, 2 - ALPHA_NUMERIC, 3 - ALPHA_NUMERIC_SYMBOL")
	fmt.Println("CIPHER_NUM_1        - Chiper number 1")
	fmt.Println("CIPHER_NUM_2        - Chiper number 2")
	fmt.Println("CIPHER_NUM_3        - Chiper number 3")
	fmt.Println("CIPHER_NUM_CHANGE   - A number used as a counter. When reached the cipher is changed.")
	fmt.Println("MESSAGE             - Message to be ecoded / decoded")
}
