package main

import "fmt"

func main() {
	Test()
}

func Test() {
	if !testCalculate("Test1", "+", 3, 2, 5) {
		return
	}

	if !testCalculate("Test1", "+", 5, 4, 9) {
		return
	}

	fmt.Println("Success!")
}

func Calculate(s string, a int, b int) int {
	return a + b
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}
