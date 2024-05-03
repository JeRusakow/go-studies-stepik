/*
Напишите элемент конвейера (функцию), что запоминает предыдущее значение и отправляет значения на следующий этап конвейера только если оно отличается от того, что пришло ранее.

Ваша функция должна принимать два канала - inputStream и outputStream, в первый вы будете получать строки, во второй вы должны отправлять значения без повторов. В итоге в outputStream должны остаться значения, которые не повторяются подряд. Не забудьте закрыть канал ;)
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	var send, recieve chan string

	send = make(chan string)
	recieve = make(chan string)

	go sender(send)
	go removeDuplicates(send, recieve)

	for r := range recieve {
		fmt.Print(r)
	}

	time.Sleep(time.Second * 10)

}

func sender(send chan string) {
	data := "aabbbccddac"
	for _, d := range data {
		send <- string(d)
	}
	close(send)
}

func removeDuplicates(inputStream, outputStream chan string) {
	var memory string

	for buf := range inputStream {
		if buf != memory {
			memory = buf
			outputStream <- buf
		}
	}

	close(outputStream)
}
