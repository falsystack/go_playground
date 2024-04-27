package main

import (
	"fmt"
	"github.com/tuckersGo/goWeb/web9/lzw"
)
import "github.com/tuckersGo/goWeb/web9/cipher"

type Component interface {
	Operator(data string)
}

var sendData string
var recvData string

type SendComponent struct {
}

func (s *SendComponent) Operator(data string) {
	// Send data
	sendData = data
}

type ZipComponent struct {
	com Component
}

func (z *ZipComponent) Operator(data string) {
	zip, err := lzw.Write([]byte(data))
	if err != nil {
		panic(err)
	}

	z.com.Operator(string(zip))
}

type EncryptComponent struct {
	com Component
	key string
}

func (e *EncryptComponent) Operator(data string) {
	encrypt, err := cipher.Encrypt([]byte(data), e.key)
	if err != nil {
		panic(err)
	}

	e.com.Operator(string(encrypt))
}

type DecryptComponent struct {
	key string
	com Component
}

func (d *DecryptComponent) Operator(data string) {
	decrypt, err := cipher.Decrypt([]byte(data), d.key)
	if err != nil {
		panic(err)
	}
	d.com.Operator(string(decrypt))
}

type UnzipComponent struct {
	com Component
}

func (u *UnzipComponent) Operator(data string) {
	unzip, err := lzw.Read([]byte(data))
	if err != nil {
		panic(err)
	}

	u.com.Operator(string(unzip))
}

type ReadComponent struct {
}

func (r ReadComponent) Operator(data string) {
	recvData = data
}

func main() {
	sender := &EncryptComponent{
		key: "abcde",
		com: &ZipComponent{
			com: &SendComponent{},
		},
	}

	sender.Operator("Hello World")
	fmt.Println(sendData)

	receiver := &UnzipComponent{
		com: &DecryptComponent{
			key: "abcde",
			com: &ReadComponent{},
		},
	}
	receiver.Operator(sendData)
	fmt.Println(recvData)
}
