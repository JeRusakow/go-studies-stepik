/*
Даны два числа. Определить цифры, входящие в запись как первого, так и второго числа.
Входные данные
Программа получает на вход два числа. Гарантируется, что цифры в числах не повторяются. Числа в пределах от 0 до 10000.
Выходные данные
Программа должна вывести цифры, которые имеются в обоих числах, через пробел. Цифры выводятся в порядке их нахождения в первом числе.
*/
package main

import "fmt"

func main() {
	var num1, num2, digit, buf uint
	fmt.Scan(&num1, &num2)
	for ; num1 > 0; num1 /= 10 {
		digit = num1 % 10

		// Searchin for the digit
		for dummy := num2; dummy > 0; dummy /= 10 {
			if dummy%10 == digit {
				// if digit found, saving it int buffer
				buf = buf*10 + digit
				break
			}
		}
	}

	// Printing the found digits backwards, as task says
	for ; buf > 0; buf /= 10 {
		fmt.Printf("%d ", buf%10)
	}
}
