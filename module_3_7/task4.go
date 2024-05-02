/*
На стандартный ввод подаются данные о продолжительности периода (формат приведен в примере). Кроме того, вам дана дата в формате Unix-Time: 1589570165 в виде константы типа int64 (наносекунды для целей преобразования равны 0).

Требуется считать данные о продолжительности периода, преобразовать их в тип Duration, а затем вывести (в формате UnixDate) дату и время, получившиеся при добавлении периода к стандартной дате.
*/

package main

import (
	"fmt"
	"time"
)

const UNIX_DATE int64 = 1589570165
const UNIX_NANOSECS int64 = 0

func main() {
	var m, s int64

	fmt.Scanf("%d мин. %d сек.", &m, &s)

	dur, _ := time.ParseDuration(fmt.Sprintf("%dm%ds", m, s))

	curDate := time.Unix(UNIX_DATE, UNIX_NANOSECS)

	curDate = curDate.Add(dur)

	fmt.Println(curDate.UTC().Format(time.UnixDate))

}
