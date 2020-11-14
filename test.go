package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

func main() {
	fmt.Println("Hello")

	l := MakeIntList(1, 2, 3, 4, 5, 6, -7, 8, 9, -10)
	var even IntPredicate = func(e int) bool { return e%2 == 0 }
	var pos IntPredicate = func(e int) bool { return e > 0 }

	l.
		Filter(even.And(pos)).
		MapString(func(e int) string { return fmt.Sprintf("~~%v~~", e) }).
		Foreach(func(e string) {
			fmt.Println(">", e)
		})

	var i1 OptionInt = Int(1)
	var i2 OptionInt = NoneInt
	fmt.Println(i1.ToString(), i1.IsDefined(), i1.IsEmpty(), i1.Filter(even).ToString(), i1.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())
	fmt.Println(i2.ToString(), i2.IsDefined(), i2.IsEmpty(), i2.Filter(even).ToString(), i2.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())

	fmt.Println(NoneInt.Equals(NoneInt))

	fmt.Println(MakeIntList(1, 2, 3, 4, 5).Size())
	fmt.Println(MakeIntList(1, 2).Size())
	fmt.Println(MakeIntList(1).Size())
	fmt.Println(NilInt.Size())

	fmt.Println(MakeIntList(1, 2, 3).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList(1, 2, 3, 4).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList().Equals(MakeIntList()))

	MakeIntList(1, 2, 3).MapInt(func(e int) int { return e * 100 }).Foreach(func(e int) { fmt.Println(">", e) })

	arr := MakeIntList(1, 2, 3, 4, 5).ToArray()
	fmt.Println(len(arr), arr[0], arr[1], arr[2], arr[3], arr[4])

	fmt.Println(MkStringIntArr([]int{1, 2, 3, 4, 5}, "~~", "|", "~~"))
}
