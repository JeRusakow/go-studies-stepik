/*
Найдите количество минимальных элементов в последовательности.
Вводится натуральное число N, а затем N целых чисел последовательности.
*/

package main

import "fmt"

func main() {
	var n, num, count, min int
	fmt.Scan(&n)

	// initialization of MIN value
	fmt.Scan(&min)
	count++

	for i := 1; i < n; i++ {
		fmt.Scan(&num)

		if num == min {
			count++
		}

		if num < min {
			min = num
			count = 1
		}

	}

	fmt.Println(count)
}
