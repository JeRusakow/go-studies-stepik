/*
Дано неотрицательное целое число. Найдите и выведите первую цифру числа.
*/

package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Printf("%c", a[0])
}
