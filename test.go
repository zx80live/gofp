package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	q := MkStringQueue("a", "b", "c").Enqueue("d").Enqueue("e").Enqueue("f")

	fmt.Println(q.Take(3).ToString(), q.TakeRight(3).ToString(), q.TakeWhile(func(e string) bool { return e != "e" }).ToString())
}
