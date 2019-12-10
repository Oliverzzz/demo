package main

import (
	"fmt"
	"net"
)

func connHandler(c net.Conn) {
	defer c.Close()
	var buf = []byte{'1','h', 'i'}
	c.Write(buf)
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1208")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	connHandler(conn)
}
