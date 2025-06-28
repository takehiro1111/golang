package main

import "fmt"

// 値型: 配列は値型です。関数に渡すときや代入するときは、配列全体がコピーされます。
func array_method() {
	var a [3]int
	a[0] = 100
	a[1] = 200
	a[2] = 300
	fmt.Println("array.go", a)
}	
