package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello " + "World")
	fmt.Println(string("Hello World"[0]))
	s1 := "Hello World"
	s2 := "Hello World2"

	s1 = strings.Replace(s1, "H", "X", 1)
	s2  = strings.Replace(s2, "H", "X", 2)
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(strings.Contains(s1, "World"))
	fmt.Println(strings.Contains("programing", "prog"))
	fmt.Println(strings.ContainsAny(s1, "word"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny(" ", " "))

	fmt.Println(`Test
                      Test
Test`)

	fmt.Println("\"")
	fmt.Println(`"`)
}
