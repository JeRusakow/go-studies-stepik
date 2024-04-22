/*
Дано трехзначное число. Найдите сумму его цифр.
*/

package main

import "fmt"

func main() {
	var num, sum int
	fmt.Scan(&num)

	for ; num > 0; num /= 10 {
		sum += num % 10
	}

	fmt.Println(sum)
}
