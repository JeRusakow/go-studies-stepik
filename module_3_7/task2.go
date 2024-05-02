/*
На стандартный ввод подается строковое представление даты и времени определенного события в следующем формате:

2020-05-15 08:00:00
Если время события до обеда (13-00), то ничего менять не нужно, достаточно вывести дату на стандартный вывод в том же формате.

Если же событие должно произойти после обеда, необходимо перенести его на то же время на следующий день, а затем вывести на стандартный вывод в том же формате.
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
		b         []byte
		eventTime time.Time
		err       error
	)

	b, _ = io.ReadAll(os.Stdin)
	s := strings.TrimSpace(string(b[:]))

	if eventTime, err = time.Parse(time.DateTime, s); err != nil {
		fmt.Println(err)
	}

	if eventTime.Hour() >= 13 {
		eventTime = eventTime.Add(time.Hour * 24)
	}

	fmt.Println(eventTime.Format(time.DateTime))
}
