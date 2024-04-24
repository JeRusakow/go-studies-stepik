/*
На вход подается строка. Нужно определить, является ли она палиндромом. Если строка является палиндромом - вывести Палиндром иначе - вывести Нет. Все входные строчки в нижнем регистре.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	runic := []rune(strings.TrimSpace(text))

	for i := 0; i < len(runic)/2; i++ {
		if runic[i] != runic[len(runic)-1-i] {
			fmt.Println("Нет")
			return
		}
	}

	fmt.Println("Палиндром")
}
