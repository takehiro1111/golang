package main

import "fmt"




func getMinNumber() {
	var minNumber int 
	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4}
	for i, num := range l {
		if i == 0 {
			minNumber = l[i]
			continue
		}

		if l[i] < l[i-1] {
			minNumber = num
		}
	}

	fmt.Println("最小値", minNumber)
}

func totalPricesFruits() {
	m := map[string]int{
    "apple":  200,
    "banana": 300,
    "grapes": 150,
    "orange": 80,
    "papaya": 500,
    "kiwi":   90,
	}

	var total int
	for _,  v := range m {
		total += v
	}

	fmt.Println("合計:", total)
}

func main() {
	getMinNumber()
	totalPricesFruits()
}
