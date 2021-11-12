package main

import "fmt"

func main() {
	// fmt.Println(Hello())
	// switchCondition()
	// deferStatement()
	// pointer()
	structStatement()
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
