/*
Дано трехзначное число. Переверните его, а затем выведите.
*/

package main

import "fmt"

func main() {
	var num, new_num int
	mult := 100

	fmt.Scan(&num)

	for ; num > 0; num /= 10 {
		new_num += (num % 10) * mult
		mult /= 10
	}

	fmt.Println(new_num)
}
