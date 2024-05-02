/*
Данная задача поможет вам разобраться в пакете encoding/csv и path/filepath, хотя для решения может быть использован также пакет archive/zip (поскольку файл с заданием предоставляется именно в этом формате).

В тестовом архиве, который вы можете скачать из нашего репозитория на github.com, содержится набор папок и файлов. Один из этих файлов является файлом с данными в формате CSV, прочие же файлы структурированных данных не содержат.

Требуется найти и прочитать этот единственный файл со структурированными данными (это таблица 10х10, разделителем является запятая), а в качестве ответа необходимо указать число, находящееся на 5 строке и 3 позиции (индексы 4 и 2 соответственно).
*/

package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func walkFunc(path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if info.IsDir() {
		return nil
	}

	if data, err := openAsCSV(path); err == nil {
		fmt.Printf("path: %s, data: %s\n", path, data)
	} else {
		return nil
	}
	return nil

}

func openAsCSV(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ','
	data, err := r.ReadAll()
	if err != nil {
		return "", err
	}
	if (len(data) == 10) && (len(data[0]) == 10) {
		return data[4][2], nil
	} else {
		return "", errors.New("no data")
	}
}

func main() {
	targetPath := ".\\task\\"
	filepath.Walk(targetPath, walkFunc)

}
