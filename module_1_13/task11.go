/*
По данному числу n закончите фразу "На лугу пасется..." одним из возможных продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова".
Дано число n (0<n<100).
*/

package main

import "fmt"

func main() {
	var a byte
	fmt.Scan(&a)

	switch {
	case a > 4 && a < 21:
		fallthrough
	case a%10 > 5 || a%10 == 0:
		fmt.Printf("%d korov", a)
	case a%10 == 1:
		fmt.Printf("%d korova", a)
	case a%10 > 1 && a%10 < 5:
		fmt.Printf("%d korovy", a)
	}
}
