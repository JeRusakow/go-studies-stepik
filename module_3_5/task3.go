/*
В тестовом файле содержится длинный ряд чисел, разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа. Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filepath := "task.data"

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	defer f.Close()

	var counter int

	r := bufio.NewReader(f)
	for str, err := r.ReadString(';'); err == nil; str, err = r.ReadString(';') {
		counter++

		if len(str) > 2 {
			continue
		}

		if str == "0;" || str == "0" {
			break
		}
	}

	fmt.Println(counter)

}
