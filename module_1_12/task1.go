/*
Напишите программу, принимающая на вход число N (N≥4), а затем N целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.
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

	fmt.Println(slice[3])
}
