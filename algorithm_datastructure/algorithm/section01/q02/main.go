package main

import (
	"fmt"
)

/*
대문자와 소문자가 같이 존재하는 문자열을 입력받아 대문자는 소문자로 소문자는 대문자로 변환하여 출력하는 프로그램을 작성하세요.

Input
StuDY

Output
sTUdy
*/

func main() {
	var str string
	fmt.Scanln(&str)

	// 65 ~ 90 -> uppercase
	// uppercase -> lowercase : uppercase + 32
	// lowercase -> uppercase : lowercase - 32

	ascii := []byte(str)
	for i := 0; i < len(ascii); i++ {
		if isUppercase(ascii, i) {
			ascii[i] = ascii[i] + 32
			continue
		}
		ascii[i] = ascii[i] - 32
	}

	fmt.Println(string(ascii))
}

func isUppercase(ascii []byte, i int) bool {
	return ascii[i] >= 65 && ascii[i] <= 90
}
