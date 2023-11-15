package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

/*
N개의 단어가 주어지면 각 단어를 뒤집어 출력하는 프로그램을 작성하세요.

Input

3
good
Time
Big

Output

doog
emiT
giB

*/

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	count, err := strconv.Atoi(sc.Text())
	if err != nil {
		return
	}

	var words []string
	for i := 0; i < count; i++ {
		sc.Scan()
		words = append(words, sc.Text())
	}

	for i := 0; i < count; i++ {
		runes := []rune(words[i])
		leftPt := 0
		rightPt := len(runes) - 1
		for leftPt < rightPt {
			runes[leftPt], runes[rightPt] = runes[rightPt], runes[leftPt]
			leftPt++
			rightPt--
		}
		fmt.Println(string(runes))
	}
}
