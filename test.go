package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l := MakeListInt(1, 2, 3, 4, 5).Cons(10).Cons(20).Cons(30)
	l.Foreach(func(e int) {
		fmt.Println(">", e)
	})
}
