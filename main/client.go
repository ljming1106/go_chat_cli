package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:2000")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connect server successfully")

	var msg string

	for {
		msg = ""
		fmt.Printf("say to server:")
		fmt.Scan(&msg)
		conn.Write([]byte(msg))
		data := make([]byte, 255)
		msg_read, err := conn.Read(data)
		if msg_read == 0 || err != nil {
			break
		}
		msg_read_str := string(data[0:msg_read])
		if msg_read_str == "close" {
			conn.Write([]byte("close"))
			break
		}

		fmt.Println("server say:", msg_read_str)
	}
	conn.Close()
}
