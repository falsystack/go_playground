package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
한 개의 문장이 주어지면 그 문장 속에서 가장 긴 단어를 출력하는 프로그램을 작성하세요.
문장속의 각 단어는 공백으로 구분됩니다.

Input
it is time to study

Output
study
*/

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	text := sc.Text()

	var idx, maxLen int

	words := strings.Split(text, " ")
	for i, word := range words {
		wordLen := len(word)
		if wordLen > maxLen {
			maxLen = wordLen
			idx = i
		}
	}

	if words[idx] == "loveispower" {
		fmt.Println("OK")
	}
}
