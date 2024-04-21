/*
Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, p float64
	var count uint
	fmt.Scan(&x, &y, &p)
	for ; ; count++ {
		if x >= p {
			break
		}
		x = math.Trunc(x * (1. + y/100.))
	}
	fmt.Println(count)
}
