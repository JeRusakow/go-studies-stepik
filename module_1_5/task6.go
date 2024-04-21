// Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m. Выведите на экран фразу:
// It is h hours m minutes.

package main

import "fmt"

func main() {
	var d, h, m uint
	fmt.Scan(&d)
	h = d / (360 / 12)
	m = d % (360 / 12) * 2
	fmt.Printf("It is %d hours %d minutes.", h, m)
}
