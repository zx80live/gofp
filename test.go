package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp/mutable"
)

func main() {
	l := NilIntLinkedList().Append(10) //.Append(20).Append(30).Append(40).Append(50)
	it := l.Iterator()
	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
