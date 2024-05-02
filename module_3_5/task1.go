/*
на стандартный ввод подаются целые числа в диапазоне 0-100, каждое число подается на стандартный ввод с новой строки (количество чисел не известно). Требуется прочитать все эти числа и вывести в стандартный вывод их сумму.
*/
package main

import (
	"bufio"
	"os"
	"strconv"
)

func main() {
	var sum int
	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for s.Scan() {
		if num, err := strconv.Atoi(s.Text()); err == nil {
			sum += num
		}
	}

	w.Write([]byte(strconv.Itoa(sum)))
	w.Flush()

}
