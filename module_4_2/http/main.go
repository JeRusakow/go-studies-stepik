/*
Считайте с консоли две переменные, сначала name, затем age. Сделайте HTTP запрос на http://127.0.0.1:8080/hello передав name и age в query параметрах, затем напечатайте ответ сервера (только тело).
*/

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var name, age string

	baseUrl := "http://127.0.0.1:8080/hello"

	fmt.Scan(&name, &age)

	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)

	fullUrl := baseUrl + "?" + params.Encode()

	resp, err := http.Get(fullUrl)
	if err != nil {
		log.Fatalln("Error sending req: ", err)
	}
	defer resp.Body.Close()

	respBuf, _ := io.ReadAll(resp.Body)
	fmt.Println(string(respBuf))

}
