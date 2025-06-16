package main

import (
	"fmt"
	"time"
	"os/user"
)

func println() {
	fmt.Println("Hello World!", time.Now())
}

func printf() {
	fmt.Printf("Hello World!\n")
	fmt.Println(user.Current())
}

func main() {
	println()
	printf()
}

