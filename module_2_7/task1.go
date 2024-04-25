/*
Дана строка, содержащая только английские буквы (большие и маленькие). Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и после последней символ ‘*’ добавлять не нужно).
*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	newString := make([]byte, 0, len(s)*2-1)

	newString = append(newString, s[0])

	for i := 1; i < len(s); i++ {
		newString = append(newString, '*', s[i])
	}

	fmt.Printf("%s", newString)

}
