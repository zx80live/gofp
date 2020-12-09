package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")
	l := MakeIntList(1, 2, 3)
	t1 := MkTuple2(10, "Hello")
	t2 := MkTuple2(20, l)
	t3 := MkTuple2(t1, t2)

	fmt.Println(l.ToString())
	fmt.Println(t1.ToString())
	fmt.Println(t2.ToString())
	fmt.Println(t3.ToString())

}
