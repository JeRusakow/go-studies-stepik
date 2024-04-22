/*
Идёт k-я секунда суток. Определите, сколько целых часов h и целых минут m прошло с начала суток.
Выведите на экран фразу:
It is ... hours ... minutes.
*/

package main

import "fmt"

func main() {
	var s int
	fmt.Scan(&s)
	fmt.Printf("It is %d hours %d minutes.", s/60/60, (s/60)%60)
}
