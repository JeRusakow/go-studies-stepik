/*
По данным числам, определите количество чисел, которые равны нулю.
Вводится натуральное число N, а затем N чисел.
*/

package main

import "fmt"

func main() {
	var n, num, count int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		if num == 0 {
			count++
		}
	}

	fmt.Println(count)
}
