/*
Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он может содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
*/

package main

import (
	"fmt"
	"unicode"
)

func main() {
	var password string
	fmt.Scan(&password)

	runes := []rune(password)

	for _, symb := range runes {
		if !(unicode.Is(unicode.Latin, symb) || unicode.Is(unicode.Digit, symb)) {
			fmt.Println("Wrong password")
			return
		}
	}

	fmt.Println("Ok")

}
