package main

import (
 "fmt"
 "net"
 "strings"
)

func main() {
    fmt.Printf("started")
    ln, er := net.Listen("tcp", ":8080")
    if er != nil {
    	fmt.Printf("error in listen")
    }
    for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Printf("error in go")
		}
    	go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Printf("error in handle")
	}
	HttpRequest := string(buf)
	requestHeaders := strings.Split(HttpRequest, "\n")
	request := requestHeaders[0]
	conn.Write([]byte(request))
}