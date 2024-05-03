/*
Внутри функции main (функцию объявлять не нужно), вам необходимо в отдельной горутине вызвать функцию work() и дождаться результатов ее выполнения.

Функция work() ничего не принимает и не возвращает.
*/

package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(time.Second * 2)
	fmt.Println("Work done")
}

func main() {
	<-func() <-chan struct{} {
		done := make(chan struct{})

		go func() {
			work()
			close(done)
		}()

		return done
	}()
}
