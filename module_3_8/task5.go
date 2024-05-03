/*
 Вам необходимо написать функцию calculator следующего вида:

func calculator(arguments <-chan int, done <-chan struct{}) <-chan int
В качестве аргумента эта функция получает два канала только для чтения, возвращает канал только для чтения.

Через канал arguments функция получит ряд чисел, а через канал done - сигнал о необходимости завершить работу. Когда сигнал о завершении работы будет получен, функция должна в выходной (возвращенный) канал отправить сумму полученных чисел.

Функция calculator должна быть неблокирующей, сразу возвращая управление.
*/

package main

import (
	"fmt"
	"time"
)

func calculator1(arguments <-chan int, done <-chan struct{}) <-chan int {
	var outputChan chan int

	outputChan = make(chan int)

	go func() {
		var sum, a int
		defer close(outputChan)

		for {
			select {
			case a = <-arguments:
				sum += a

			case <-done:
				outputChan <- sum
				return
			}
		}

	}()

	return outputChan
}

func main() {
	args := make(chan int)
	stop := make(chan struct{})

	defer close(args)
	defer close(stop)

	nums := []int{1, 2}

	out := calculator1(args, stop)

	// test it here
	for _, n := range nums {
		args <- n
	}
	time.Sleep(time.Second)
	stop <- struct{}{}

	fmt.Println(<-out)

}
