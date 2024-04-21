package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

/*
*
01_writeを実行
06_dial-readを実行
*/
func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	bs, err := io.ReadAll(conn)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(bs))
}
