/*
На вход подается строка. Нужно определить, является ли она правильной или нет. Правильная строка начинается с заглавной буквы и заканчивается точкой. Если строка правильная - вывести Right иначе - вывести Wrong
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	runic := []rune(strings.TrimSpace(text))

	if unicode.IsUpper(runic[0]) && (runic[len(runic)-1] == rune('.')) {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}

}
