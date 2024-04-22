/*
Даны три натуральных числа a, b, c. Определите, существует ли треугольник с такими сторонами.
*/

package main

import "fmt"

func main() {
	var a, b, c, p int
	fmt.Scan(&a, &b, &c)

	p = (a + b + c) / 2
	if (p-a)*(p-b)*(p-c) > 0 {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}
