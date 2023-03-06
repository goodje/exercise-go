package closure

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	i := 0
	a := 0
	b := 1
	d := func() {
		i = i + 1
	}
	return func() int {
		defer d()
		if i == 0 {
			return a
		} else if i == 1 {
			return b
		}

		v := a + b
		a = b
		b = v
		return v
	}
}

func main1() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
