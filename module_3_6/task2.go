/*
мы будем использовать справочник ОКВЭД (Общероссийский классификатор видов экономической деятельности), который был размещен на web-портале data.gov.ru.

Необходимая вам информация о структуре данных содержится в файле structure-20190514T0000.json, а сами данные, которые вам потребуется декодировать, содержатся в файле data-20190514T0100.json. Файлы размещены в нашем репозитории на github.com.

Для того, чтобы показать, что вы действительно смогли декодировать документ вам необходимо в качестве ответа записать сумму полей global_id всех элементов, закодированных в файле.
*/

package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Entry struct {
	Global_Id      int64  `json: "global_id"`
	SystemObjectID string `json: "system_object_id"`
	SignatureDate  string `json: "signature_date"`
	Razdel         string `json: "Razdel"`
	Kod            string `json: "Kod"`
	Name           string `json: "Name"`
	Idx            string `json: "Idx"`
}

type OKVED struct {
	Entries []Entry
}

func main() {
	var okved []Entry
	var sum int64

	filepath := "data-20190514T0100.json"
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	bytes, _ := io.ReadAll(f)

	err = json.Unmarshal(bytes, &okved)
	if err != nil {
		fmt.Println(err)
	}

	for _, entry := range okved {
		sum += entry.Global_Id
	}

	fmt.Println(sum)
}
