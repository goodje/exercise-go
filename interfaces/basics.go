package main

import "fmt"

type I interface {
	met(value string)
}

type T struct {
	S string
}

func (t *T) met(value string) {
	if value != "" {
		t.S = value
	}

	fmt.Println(t.S)
}

func main() {
	// assertion
	var ia interface{}
	ia = "hello"

	sa := ia.(string)
	fmt.Println(sa)
	sa1, ok := ia.(int)
	fmt.Println(sa1, ok)

	// interface
	t := T{"title"}
	t.met("")

	// assign an interface value to another one
	var i I
	var i2 I
	i = &T{"title"}
	i2 = i

	i2.met("i2 met")
	i.met("")
}
