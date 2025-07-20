package main

import "fmt"

func one(x *int) {
	fmt.Println("メモリ:",x)
	*x=1
}

// func one(x *int) {
// 	fmt.Println("メモリ:",&x)
// 	*x=1
// }

func main() {
	var n int = 100
	one(&n)
	fmt.Println(n)
	fmt.Println(&n)
}

