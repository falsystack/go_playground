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
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		return 0
	}
}

func testCalculate(testcase, op string, a, b, expected int) bool {
	o := Calculate(op, a, b)
	if o != expected {
		fmt.Printf("%s Failed! expected:%d output:%d\n", testcase, expected, o)
		return false
	}
	return true
}
