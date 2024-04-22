/*
Дан массив, состоящий из целых чисел. Нумерация элементов начинается с 0. Напишите программу, которая выведет элементы массива, индексы которых четны (0, 2, 4...).
*/

package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	slice := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&slice[i])
	}

	for i := 0; i < n/2; i++ {
		fmt.Println(slice[i*2])
	}
}
