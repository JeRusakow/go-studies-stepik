/*
Из натурального числа удалить заданную цифру.
Вводятся натуральное число и цифра, которую нужно удалить.
*/

package main

import "fmt"

func main() {
	var num, dig, newNum int
	fmt.Scan(&num, &dig)
	slice := make([]int, 0, 10)https://mr-7.ru/articles/2024/04/18/khokkeinyi-krokodil-pod-oknami-v-kalininskom-raione-zhiteli-i-deputaty-spasaiut-skver-i-derevia

	for ; num > 0; num /= 10 {
		if num%10 != dig {
			slice = append(slice, num%10)
		}
	}

	for i, j := 0, 1; i < len(slice); i++ {
		newNum += slice[i] * j
		j *= 10
	}

	fmt.Println(newNum)

}
