/*
Последовательность состоит из натуральных чисел и завершается числом 0. Определите количество элементов этой последовательности, которые равны ее наибольшему элементу.
*/
package main

import "fmt"

func main() {
	var num, max, count int

	for fmt.Scan(&num); ; fmt.Scan(&num) {
		if num > max {
			max = num
			count = 0
		}

		if num == max {
			count++
		}

		if num == 0 {
			break
		}
	}
	fmt.Println(count)
}
