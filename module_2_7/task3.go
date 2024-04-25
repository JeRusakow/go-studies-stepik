/*
На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.

Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
*/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	str2 := make([]byte, 0, len(s)*2)

	for i := 0; i < len(s); i++ {
		str2 = strconv.AppendUint(str2, uint64((s[i]-'0')*(s[i]-'0')), 10)
	}

	fmt.Printf("%s", str2)

}
