package main

import "fmt"

func main() {
	// fmt.Println(Hello())
	// switchCondition()
	// deferStatement()
	// pointer()
	// structStatement()
	arraysnslices()
}

func arraysnslices() {
	// arrays can't be resized
	var a [7]int = [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println(a)

	var s []int = a[3:5]
	fmt.Println(s) // 7, 11
	s[0] = 8       // changes the underlying array
	fmt.Println(a)

	// create a slice, it automatically creates an array first
	s2 := []int{5, 6, 7}
	fmt.Println(s2)

	s3 := s2[:]
	fmt.Println(s3)
}

// struct
type Cat struct {
	name      string
	bodyColor string
}

func structStatement() {
	cat := Cat{"lulu", "yellow"}
	fmt.Println("cat name:", cat.name)
	fmt.Println("cat body color:", cat.bodyColor)

	p := &cat
	fmt.Println("*pcat name:", (*p).name)
	fmt.Println("pcat name:", p.name)
}

// end struct

func pointer() {
	i := 5
	p := &i
	fmt.Println(p, *p)
}

func deferStatement() {
	defer fmt.Println("defer print")

	for i := 1; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("return print")
}

func switchCondition() {
	switch {
	case 5 < 0:
		fmt.Println("5 < 0")
	case 5 > 0:
		fmt.Println("5 > 0")
	default:
		fmt.Print("not sure")
	}
}

func forloop() {
	i := 5
	for {
		if i < 10 {
			fmt.Println(i)
			i++
			continue
		}
		break
	}
}

const hello_world = "Hello, World"

func Hello() string {
	return hello_world
}
