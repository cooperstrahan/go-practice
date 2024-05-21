package main

import (
	"bufio"
	"fmt"
	"net"
)

func connect_tcp() {
	// simple tcp connection
	conn, _ := net.Dial("tcp", "golang.org:80")
	count, _ := fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	fmt.Println(count)
	status, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("printing status")
	fmt.Println(status)
}