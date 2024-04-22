/*
Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи оно является, то есть выведите такое число n, что φ(n)=A. Если А не является числом Фибоначчи, выведите число -1.
*/

package main

import "fmt"

func main() {
	var a int
	p, q := 0, 1
	fmt.Scan(&a)

	if a == p {
		fmt.Println(0)
		return
	}

	if a == q {
		fmt.Println(1)
		return
	}

	for i := 2; ; i++ {
		p, q = q, p+q
		if a == q {
			fmt.Println(i)
			return
		}

		if a < q {
			fmt.Println(-1)
			return
		}
	}
}
