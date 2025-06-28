package main

import "fmt"

func slice_method() {
	n := [] int{1,2,3,4,5,6}

	fmt.Println(n)
	fmt.Println(n[2])
	fmt.Println(n[2:4])
	fmt.Println(n[2:])
	fmt.Println(n[:])

	// スライスの末尾に要素を追加する。
	n = append(n, 100, 200, 300, 400)

	fmt.Println(n)

	var board = [][] int {
		{0,1,2},
		{3,4,5},
		{6,7,8},
	}

	fmt.Println(board)

	
	fmt.Println(n)
}


