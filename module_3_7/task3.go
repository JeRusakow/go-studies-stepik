/*
На стандартный ввод подается строковое представление двух дат, разделенных запятой (формат данных смотрите в примере).

Необходимо преобразовать полученные данные в тип Time, а затем вывести продолжительность периода между меньшей и большей датами.
*/

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func main() {
	var (
		times [2]time.Time
	)

	b, _ := io.ReadAll(os.Stdin)
	s := strings.Split(strings.TrimSpace(string(b[:])), ",")

	for idx, str := range s {
		times[idx], _ = time.Parse("02.01.2006 15:04:05", str)
	}

	dur := times[0].Sub(times[1])

	if dur < 0 {
		dur = -dur
	}

	fmt.Println(dur)
}
