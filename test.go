package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	l := MkIntListQueue(MkIntList(1, 2, 3), MkIntList(), MkIntList(4, 5, 6))
	it := l.FlatMapInt(func(e IntList) IntQueue {
		inner := e.Iterator()
		q := MkIntQueue()
		for inner.HasNext() {
			q = q.Enqueue(inner.Next())
		}
		return q
	}).Iterator()

	for it.HasNext() {
		fmt.Println(it.Next())
	}
}
