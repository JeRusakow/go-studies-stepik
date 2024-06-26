/*
Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.
Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .
По данному числу определите его цифровой корень.
*/

package main

import "fmt"

func main() {
	var a int64
	fmt.Scan(&a)
	fmt.Println(digital_root(a))
}

func digital_root(a int64) int64 {
	if a < 10 {
		return a
	}
	var sum int64
	for ; a > 0; a /= 10 {
		sum += a % 10
	}
	return digital_root(sum)

}
