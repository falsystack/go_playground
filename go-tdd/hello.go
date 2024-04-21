package main

import "fmt"

var opMap map[string]func(int, int) int

func main() {
	initialOpMap()
	Test()
}

// use strategy pattern
func initialOpMap() {
	opMap = make(map[string]func(int, int) int)
	opMap["+"] = add
	opMap["-"] = sub
	opMap["*"] = mul
	opMap["/"] = div
}

func div(a int, b int) int {
	return a / b
}

func mul(a int, b int) int {
	return a * b
}

func sub(a int, b int) int {
	return a - b
}

func add(a int, b int) int {
	return a + b
}

func Test() {
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	if !testCalculate("Test1", "+", 5, 4, 9) {
		return
	}

	if !testCalculate("Test3", "-", 5, 3, 2) {
		return
	}

	if !testCalculate("Test4", "-", 3, 6, -3) {
		return
	}

	if !testCalculate("Test5", "*", 3, 7, 21) {
		return
	}

	if !testCalculate("Test6", "*", 3, 0, 0) {
		return
	}

	if !testCalculate("Test7", "*", 3, -3, -9) {
		return
	}

	if !testCalculate("Test8", "/", 9, 3, 3) {
		return
	}

	fmt.Println("Success!")
}

func Calculate(op string, a int, b int) int {
	if f, ok := opMap[op]; ok {
		return f(a, b)
	}
	return 0
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}
