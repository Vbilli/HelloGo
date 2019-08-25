package main

import (
	"fmt"
	"runtime"
)

// Vertex should be the one
type Vertex struct {
	x int
	y int
}

func main() {
	fmt.Println(add(43, 23))
	fmt.Println(swap("wait", "world"))
	fmt.Println(split(17))
	var i int
	var c, python, java bool
	fmt.Println(i, c, python, java)
	getCase1()
	getPointer()
	fmt.Println(Vertex{1, 3})

	// slice
	p := []int{2, 3, 5, 6, 11, 13}
	fmt.Println("p ==", p)

	for y := 0; y < len(p); y++ {
		fmt.Println("p[%d] == %d\n", y, p[y])
	}

	getConst()
	// make slice

	// cld := make([]int, 5)
	// dd := &cld
	// fmt.Println(dd)
}

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func getCase1() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.", os)

	}
}

func getConst() {
	const (
		a = iota
		b
		c
		d = "ha"
		e
		f = iota
		g = 1
	)
	fmt.Println(a, b, c, d, e, f, g)
}

func getPointer() {
	i, j := 42, 270
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)
}
