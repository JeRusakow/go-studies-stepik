/*
Даны два числа. Найти их среднее арифметическое.
*/

package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)

	fmt.Println((a + b) / 2.)
}
