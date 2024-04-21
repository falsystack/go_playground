package main

import (
	"fmt"
	"log"
	"net"
)

/*
02_readを実行
06_dial-writeを実行
*/

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	fmt.Fprintln(conn, "I dialed you.")
}
