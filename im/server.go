package main

import (
	"fmt"
	"net"
	"strconv"
)

func connHandler(c net.Conn, users *map[string]net.Conn) {
	buf := make([]byte, 1024)
	for {
		cnt, err := c.Read(buf)
		if err != nil || cnt ==0{
			c.Close()
			break
		}
		des := string(buf[0])
		receiconn := (*users)[des]
		receiconn.Write(buf[1:cnt])
		fmt.Println("1")
	}
}

func main() {
	server, _ := net.Listen("tcp", ":1208")
	var id = 1
	users := map[string]net.Conn{}
	for {
		conn, _ := server.Accept()
		users[strconv.Itoa(id)]  = conn
		id +=1
		go connHandler(conn, &users)
	}
}
