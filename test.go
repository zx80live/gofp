package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l := MakeListInt(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	l.
		Filter(func(e int) bool { return e%2 == 0 }).
		MapString(func(e int) string { return fmt.Sprintf("~~%v~~", e) }).
		Foreach(func(e string) {
			fmt.Println(">", e)
		})
}
