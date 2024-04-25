/*
Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
*/

package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)

	max := s[0]
	for i := 1; i < len(s); i++ {
		if s[i] > max {
			max = s[i]
			if max == '9' {
				break
			}
		}
	}

	fmt.Printf("%c", max)
}
