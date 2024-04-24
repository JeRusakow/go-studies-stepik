/*
Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	idx := 0

	for {
		if idx >= utf8.RuneCountInString(text) {
			break
		}
		if strings.Index(text[idx+1:], (text[idx:idx+1])) > -1 {
			text = strings.ReplaceAll(text, text[idx:idx+1], "")
			idx--
		}
		idx++
	}

	fmt.Printf("%s", text)
}
