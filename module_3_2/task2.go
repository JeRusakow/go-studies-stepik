/*
В привычных нам редакторах электронных таблиц присутствует удобное представление числа с разделителем разрядов в виде пробела, кроме того в России целая часть от дробной отделяется запятой. Набор таких чисел был экспортирован в формат CSV, где в качестве разделителя используется символ ";".
На стандартный ввод вы получаете 2 таких вещественных числа, в качестве результата требуется вывести частное от деления первого числа на второе с точностью до четырех знаков после "запятой" (на самом деле после точки, результат не требуется приводить к исходному формату).
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toStandartFloatFormatiing(s string) (float64, error) {
	s = strings.Replace(s, " ", "", -1)
	s = strings.Replace(s, ",", ".", 1)
	res, err := strconv.ParseFloat(s, 64)
	return res, err
}

func main() {
	str, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	str = strings.TrimSpace(str)
	str1, str2 := make([]byte, 0, 0), make([]byte, 0, 0)

	semicolumnIdx := strings.Index(str, ";")

	str1 = append(str1, str[:semicolumnIdx]...)
	str2 = append(str2, str[semicolumnIdx+1:]...)

	a, _ := toStandartFloatFormatiing(string(str1))
	b, _ := toStandartFloatFormatiing(string(str2))

	fmt.Printf("%.4f", a/b)

}
