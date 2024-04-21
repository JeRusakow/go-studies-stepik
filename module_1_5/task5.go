// Дано неотрицательное целое число. Найдите число десятков (то есть вторую цифру справа).

package main

import "fmt"

func main() {
	var a uint
	fmt.Scan(&a)
	fmt.Println(a / 10 % 10)
}
