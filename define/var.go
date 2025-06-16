package main

import (
	"fmt"
)

var (
	i    int     = 1
	f64  float64 = 1.2
	s    string  = "test"
	t, f bool    = true, false
)

func foo() {
	xi := 1
	var xf32 float32 = 1.2
	var x64 string = "x64"
	xs := "test"
	xt, xf := true, false
	fmt.Println(xi, xf32,x64, xs, xt, xf)
	fmt.Printf("%T\n", xf32)
	fmt.Printf("%T\n", xi)
}

