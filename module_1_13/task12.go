/*
По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	for i := 1; i <= n; i *= 2 {
		fmt.Printf("%d ", i)
	}
}
