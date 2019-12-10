package main

import (
	"fmt"
	"net"
)

func connHandler(c net.Conn) {
	defer c.Close()
	var buf = make([]byte, 1024)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt == 0 {
			c.Close()
			break
		}
		inStr := string(buf[0:cnt])
		fmt.Println(inStr)
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:1208")
	if err != nil {
		fmt.Printf("Fail to connect, %s\n", err)
		return
	}
	connHandler(conn)
}
