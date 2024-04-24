/*
 На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	runic := []rune(text)

	newString := make([]rune, 0, len(runic)/2)

	for i := 0; i < len(runic)/2; i++ {
		newString = append(newString, runic[i*2+1])
	}

	fmt.Printf("%s", string(newString))
}
