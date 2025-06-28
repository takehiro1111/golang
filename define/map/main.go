package main

import (
	"fmt"
)

func main() {
	m := map[string] int { "apple": 1, "banana": 2, "test": 3}

	fmt.Println(m)
	fmt.Println(m["banana"])

	v, ok := m["apple"]
	fmt.Println(v, ok)

	v, no := m["test2"]
	fmt.Println(v, no)

	m["cdd"] = 100
	fmt.Println(m)

	m3 := make(map[string] int)
	m3["pc"] = 5000
	fmt.Println(m3)

	var nillList []int
	if nillList == nil {
		fmt.Println("nilList")
	}

	var nillMap map[string] int
	if nillMap == nil {
		fmt.Println("nilMap")
	}

}
