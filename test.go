package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	q := MkStringQueue("a", "b", "c").Enqueue("d").Enqueue("e").Enqueue("f")
	q.Foreach(func(e string) { fmt.Println(e) })
	fmt.Println(q.ToList().ToString())

	q2 := MkIntQueue(1, 2, 3, 4, 5, 6, 7).Enqueue(8).Enqueue(9)
	fmt.Println(q2.Filter(EvenInt).ToList().ToString())

	fmt.Println(q.Reduce(func(a, b string) string { return fmt.Sprintf("%v-%v", a, b) }))
	fmt.Println(q2.Reduce(func(a, b int) int { return a + b }))

	fmt.Println(q.MapString(func(e string) string { return fmt.Sprintf("<%v>", e) }).ToList().ToString())

	fmt.Println(q2.MapString(func(e int) string { return fmt.Sprintf("<%v>", e) }).ToList().ToString())

}
