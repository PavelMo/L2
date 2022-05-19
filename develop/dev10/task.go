package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

/*
=== Утилита telnet ===
Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123
Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	host := flag.String("host", "localhost", "Хост")
	port := flag.String("port", "8080", "Порт")
	timeOut := flag.Duration("timeout", 10*time.Second, "Таймаут")
	flag.Parse()
	address := (*host) + ":" + (*port)
	//Запускаем сервер в отдельной горутине
	go runServer(address)

	time.Sleep(time.Second)

	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	go writeToSocket(conn)
	for {
		err = conn.SetDeadline(time.Now().Add(*timeOut))
		if err != nil {
			log.Fatalln(err)
		}
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				log.Println("timeout:", err)
			} else {
				log.Fatalln(err)
			}
			continue
		}
		fmt.Println("Get message from the socket: ", text)
	}

}
func writeToSocket(conn net.Conn) {
	defer closeConnection(conn)
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadBytes('\n')
		if err == io.EOF {
			fmt.Println("Got ctrl+D signal.Closing connection")
			return
		}
		_, err = conn.Write(text)
		if err != nil {
			fmt.Println("Connection closed due to error:", err)
			return
		}
	}
}
func runServer(address string) {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server is listening on:", address)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			closeConnection(conn)
			os.Exit(1)
		}
		text, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
			return
		}
		//fmt.Println("Server get message: ", text)
		_, err = conn.Write([]byte("New server answer: " + text + "\n"))
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}
func closeConnection(conn net.Conn) {
	err := conn.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
