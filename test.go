package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l := MakeListInt(1, 2, 3, 4, 5)
	fmt.Println(l.Head())
}
