package main

import (
	"fmt"
	"strings"
)

/*
한 개의 문자열을 입력받고, 특정 문자를 입력받아 해당 특정문자가 입력받은 문자열에 몇 개 존재하는지 알아내는 프로그램을 작성하세요.

Input
Computercooler
c

Output
2
*/
func main() {
	var inputText, target string
	var count = 0

	fmt.Scanln(&inputText)
	fmt.Scanln(&target)

	words := strings.Split(inputText, "")
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == target {
			count++
		}
	}

	fmt.Println(count)
}
