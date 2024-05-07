/*
Подключитесь к адресу 127.0.0.1:8081 по протоколу TCP, считайте от сервера 3 сообщения, и выведите их в верхнем регистре. Рекомендуем использовать буфер в 1024 байта.
*/

package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func server() {

	msg := "MixEdCaSe MesSAge"

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: 8081})
	if err != nil {
		log.Fatalln("Server, Setting listener: ", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatalln("Server, Opening connection: ", err)
	}
	defer conn.Close()

	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 2)
		conn.Write([]byte(msg))
	}
}

func main() {
	go server()

	buf := make([]byte, 1024)

	conn, err := net.DialTCP("tcp", nil, &net.TCPAddr{IP: []byte{127, 0, 0, 1}, Port: 8081})
	if err != nil {
		log.Fatalln("Setting connection: ", err)
	}
	defer conn.Close()

	for i := 0; i < 3; i++ {
		n, err := conn.Read(buf)
		if err != nil {
			log.Fatalln("Reading from conn: ", err)
		}

		fmt.Print(strings.ToUpper(string(buf[:n])))
	}

}
