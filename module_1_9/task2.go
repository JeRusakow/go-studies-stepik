/*
По данному трехзначному числу определите, все ли его цифры различны.
*/

package main

import "fmt"

func main() {
	var a string
	fmt.Scan(&a)

	if (a[0] != a[1]) && (a[1] != a[2]) && (a[0] != a[2]) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
