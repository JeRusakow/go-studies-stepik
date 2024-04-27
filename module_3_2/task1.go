/*
Вы пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.

Но не было бы так все просто, ведь ваш коллега решил подшутить над вами. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
Поэтому предварительно вам нужно убрать из них мусор и конвертировать в числа. Под мусором имеются ввиду лишние символы и спец знаки. Разрешается использовать только определенные пакеты: fmt, strconv, unicode, strings,  bytes. Они уже импортированы,
*/

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func prepareString(a []rune) []rune {
	res := make([]rune, 0, len(a))
	for _, char := range a {
		if unicode.IsDigit(char) {
			res = append(res, char)
		}
	}

	if len(res) == 0 {
		res = append(res, '0')
	}
	return res
}

func adding(a, b string) int64 {
	aClean := prepareString([]rune(a))
	bClean := prepareString([]rune(b))

	aParsed, _ := strconv.ParseInt(string(aClean), 10, 64)
	bParsed, _ := strconv.ParseInt(string(bClean), 10, 64)
	return aParsed + bParsed
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)
	fmt.Println(adding(a, b))
}
