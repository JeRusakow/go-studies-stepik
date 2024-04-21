/*
Определите является ли билет счастливым. Счастливым считается билет, в шестизначном номере которого сумма первых трёх цифр совпадает с суммой трёх последних.
*/

package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	if a[0]+a[1]+a[2] == a[3]+a[4]+a[5] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
