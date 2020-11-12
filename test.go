package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l := MakeListInt(1, 2, 3, 4, 5, 6, -7, 8, 9, -10)
	var even PredicateInt = func(e int) bool { return e%2 == 0 }
	var pos PredicateInt = func(e int) bool { return e > 0 }

	l.
		Filter(even.And(pos)).
		MapString(func(e int) string { return fmt.Sprintf("~~%v~~", e) }).
		Foreach(func(e string) {
			fmt.Println(">", e)
		})

	var i1 OptionInt = Int(1)
	var i2 OptionInt = NoneInt
	fmt.Println(i1.ToString(), i1.IsDefined(), i1.IsEmpty())
	fmt.Println(i2.ToString(), i2.IsDefined(), i2.IsEmpty())
}
