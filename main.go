package main

import (
	"fmt"
	"time"

	"github.com/heimbogdan/go-enigma/internal/cipher"
)

func main() {
	start := time.Now()
	cipher.SetScramble(cipher.ALPHA_NUMERIC_SYMBOLS, 11, 25, 3, 10)
	encMessage := cipher.Encode("In acest \"jurnal de bord\" nu o sa scriu numai despre consum. O sa scriu si despre alte opinii personale pe masura ce ma voi putea pronunta in acest sens. Nu pot spune ceva legat de cum se aseaza la drum sau cum \"trage\". Am condus-o doar in Bucuresti pana in prezent. Are doar 238 km la bord in acest moment.")
	fmt.Println(encMessage)
	fmt.Println(cipher.Decode(encMessage))
	end := time.Now()
	elapsed := end.Sub(start)
	fmt.Printf("Total duration: %d microseconds", elapsed.Microseconds())
}
