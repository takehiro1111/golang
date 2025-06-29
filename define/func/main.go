package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func add2(num1 ,num2 int) int {
	return num1 + num2
}

func add3(num1, num2 int) (result int) {
	result = num1 + num2
	return
}

func calc(num1 ,num2, num3 int) int {
	return (num1 + num2 + num3) / 3
}

func calcAdd(num1 ,num2, num3 int) (total int, calc int ) {
	total = num1 + num2 + num3
	calc = total / 3

	if calc < 0 && total < 0 {
		return 0, 0
	}
	return calc, total
}


func main() {
	fmt.Println(add(1,2))
	fmt.Println(add2(1,2))
	fmt.Println(add3(1,2))
	fmt.Println(calc(1, 2, 3))
	testFunc1, testFunc2  := calcAdd(1, 2, 3)
	fmt.Println(testFunc1, testFunc2)
}
