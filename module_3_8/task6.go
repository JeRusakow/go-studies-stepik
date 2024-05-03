/*
Необходимо написать функцию func merge2Channels(fn func(int) int, in1 <-chan int, in2 <- chan int, out chan<- int, n int).

Описание ее работы:
n раз сделать следующее:
-прочитать по одному числу из каждого из двух каналов in1 и in2, назовем их x1 и x2.
-вычислить f(x1) + f(x2)
-записать полученное значение в out

Функция merge2Channels должна быть неблокирующей, сразу возвращая управление.

Функция fn может работать долгое время, ожидая чего-либо или производя вычисления.

Формат ввода:

количество итераций передается через аргумент n.
целые числа подаются через аргументы-каналы in1 и in2.
функция для обработки чисел перед сложением передается через аргумент fn.
Формат вывода:

канал для вывода результатов передается через аргумент out.
*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func merge2Channels(fn func(int) int, in1 <-chan int, in2 <-chan int, out chan<- int, n int) {

	go func() {

		var fnStorage map[int]int
		var sentArguments []int
		var wg sync.WaitGroup
		var mut sync.Mutex
		var count1, count2 int

		sentArguments = make([]int, 2*n)
		fnStorage = make(map[int]int)

		var arg int

		for i := 0; i < 2*n; i++ {

			select {
			case arg = <-in1:
				sentArguments[count1*2] = arg
				// fmt.Printf("Read %d from In1\n", arg)
				count1++
			case arg = <-in2:
				sentArguments[count2*2+1] = arg
				// fmt.Printf("Read %d from In2\n", arg)
				count2++
			}

			if _, ok := fnStorage[arg]; !ok {
				wg.Add(1)
				go func(arg int) {
					res := fn(arg)
					mut.Lock()
					fnStorage[arg] = res
					mut.Unlock()
					wg.Done()
				}(arg)
			}
		}

		wg.Wait()
		// fmt.Println(sentArguments)

		for i := 0; i < n; i++ {
			// fmt.Printf("storage[%d] = %d  +  storage[%d] = %d   =  %d\n", sentArguments[i*2], fnStorage[sentArguments[i*2]], sentArguments[i*2+1], fnStorage[sentArguments[i*2+1]], fnStorage[sentArguments[i*2]]+fnStorage[sentArguments[i*2+1]])
			out <- fnStorage[sentArguments[i*2]] + fnStorage[sentArguments[i*2+1]]
		}
	}()
}

func blop(arg int) int {
	time.Sleep(time.Second * 2)
	return arg
}

func main() {

	const N int = 10

	in1 := make(chan int)
	in2 := make(chan int)
	out := make(chan int)

	merge2Channels(blop, in1, in2, out, N)

	go func() {
		for i := 0; i < N; i++ {
			in1 <- i
		}
	}()

	go func() {
		for i := 0; i < N; i++ {
			in2 <- i
		}
	}()

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", <-out)
		// <-out
	}

}
