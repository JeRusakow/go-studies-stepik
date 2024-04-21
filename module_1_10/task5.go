/*
Найдите первое число от 1 до n включительно, кратное c, но НЕ кратное d.
Вводится 3 натуральных числа n, c, d, каждое из которых не превышает 10000.
*/

package main

import "fmt"

func main() {
	var n, c, d int
	fmt.Scan(&n, &c, &d)

	for i := 1; i <= n; i++ {
		if (i%c == 0) && (i%d != 0) {
			fmt.Println(i)
			break
		}
	}
}
