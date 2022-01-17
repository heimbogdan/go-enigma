package main

import (
	"fmt"
	"time"

	"github.com/heimbogdan/go-enigma/cipher"
)

func main() {
	start := time.Now()
	cipher.SetScramble(cipher.ALPHA_NUMERIC_SYMBOLS, 11, 25, 3, 3)
	encMessage := cipher.Encode("Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.")
	fmt.Println(encMessage)
	fmt.Println(cipher.Decode(encMessage))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total duration: %d microseconds", elapsed.Microseconds())
}
