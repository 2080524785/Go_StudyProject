package main

import (
	"fmt"
	"math"
)

func main() {
	study4()
}
func study1() {
	fmt.Println("hello world")
}

// 变量
func study2() {
	var a = "initial"

	var b, c int = 1, 2

	var d = true

	var e float64

	f := float32(e)

	g := a + "foo"
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g)

	const s string = "constant"
	const h = 50000
	const i = 3e20 / h
	fmt.Println(s, h, i, math.Sin(h), math.Sin(i))
}

// if else
func study3() {
	if 7%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
	if num := 9; num < 0 {
		fmt.Println("1")
	} else if num < 10 {
		fmt.Println("2")
	} else {
		fmt.Println("3")
	}
}

// 循环
func study4() {
	i := 1
	for {
		fmt.Println("1")
		break
	}
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	for n := 0; n < 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("2")
	}
	for i <= 3 {
		fmt.Println(i)
		i++
	}
}
