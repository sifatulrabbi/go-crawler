package main

import (
	"fmt"
	"log"
	"net"
)

const (
	KB = 1024
	MB = 1024 * KB
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:9001")
	failOnErr(err)

	tcpListener, err := net.ListenTCP("tcp", addr)
	failOnErr(err, "main(): Listening to port.")
	defer tcpListener.Close()

	fmt.Println("started the tcp listener", tcpListener.Addr())
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn *net.TCPConn) {
	defer conn.Close()
	buf := make([]byte, 512*MB)
	if _, err := conn.Read(buf); err != nil {
		fmt.Println("error while reading:", err)
	}
	fmt.Printf("---body start---\n%s\n---body end---\n\n", buf)
}

func failOnErr(err error, msg ...string) {
	if err != nil {
		log.Fatalln(nil, msg)
	}
}
