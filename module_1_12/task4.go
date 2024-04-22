/*
Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.
*/
package main

import "fmt"

func main() {
	var n, count int
	fmt.Scan(&n)
	slice := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < n; i++ {
		if slice[i] > 0 {
			count++
		}
	}
	fmt.Print(count)
}
