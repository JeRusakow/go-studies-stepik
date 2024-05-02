/*
На стандартный ввод подаются данные о студентах университетской группы в формате JSON:

{
    "ID":134,
    "Number":"ИЛМ-1274",
    "Year":2,
    "Students":[
        {
            "LastName":"Вещий",
            "FirstName":"Лифон",
            "MiddleName":"Вениаминович",
            "Birthday":"4апреля1970года",
            "Address":"632432,г.Тобольск,ул.Киевская,дом6,квартира23",
            "Phone":"+7(948)709-47-24",
            "Rating":[1,2,3]
        },
        {
            // ...
        }
    ]
}
В сведениях о каждом студенте содержится информация о полученных им оценках (Rating). Требуется прочитать данные, и рассчитать среднее количество оценок, полученное студентами группы. Ответ на задачу требуется записать на стандартный вывод в формате JSON в следующей форме:

{
    "Average": 14.1
}

*/

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	LastName, FirstName, MiddleName, Birthday, Address, Phone string
	Rating                                                    []int
}

type Group struct {
	ID, Year int
	Number   string
	Students []Student
}
type Average struct {
	Average float32
}

func main() {
	var (
		group   Group
		average Average
		lenSum  int
	)

	json.NewDecoder(os.Stdin).Decode(&group)

	for _, student := range group.Students {
		lenSum += len(student.Rating)
	}

	average.Average = float32(lenSum) / (float32(len(group.Students)))

	data, _ := json.MarshalIndent(average, "", "    ")
	fmt.Printf("%s", data)
}
