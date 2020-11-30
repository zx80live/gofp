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

	var i1 IntOption = Int(1)
	var i2 IntOption = NoneInt

	fmt.Println(i1.ToString(), i1.IsDefined(), i1.IsEmpty(), i1.Filter(even).ToString(), i1.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())
	fmt.Println(i2.ToString(), i2.IsDefined(), i2.IsEmpty(), i2.Filter(even).ToString(), i2.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())

	fmt.Println(NoneIntOption.Equals(NoneIntOption))

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

	fmt.Println(IntArrayMkString([]int{1, 2, 3, 4, 5}, "~~", "|", "~~"))
	nestedList := MakeIntArrayList([]int{1, 2, 3}, []int{4, 5, 6}, []int{7})
	fmt.Println(nestedList.MkString("[", ",", "]"))
	fmt.Println(nestedList.ToString())

	var x int = 100
	var isXdefined bool = &x == nil
	fmt.Println(isXdefined)

	mapString := MakeIntList(1, 2, 3, 4, 5).MapString(func(e int) string { return fmt.Sprintf("{%v}", e) })
	mapString.Foreach(func(e string) {
		fmt.Println("> ", e)
	})
	fmt.Println(mapString.ToString())

	fmt.Println(Int(10).ToString())

	fmt.Println(MakeIntList(1, 2).Tail().Tail().HeadOption().ToString())

	fmt.Println(MakeIntList(1, 2, 3).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList(1, 2).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5)).
		Equals(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5))))

	MakeIntListList(MakeIntList(1, 2, 3), MakeIntList(4, 5), MakeIntList(), MakeIntList(6, 7, 8)).
		FlatMapInt(func(l IntList) IntList { return l }).
		Foreach(func(e int) { fmt.Println(">", e) })

	MakeIntOptionOption(Int(10)).FlatMapInt(func(e IntOption) IntOption { return e }).Foreach(func(e int) { fmt.Println("flatten", e) })

	MakeIntListList(MakeIntList(1, 2, 3), MakeIntList(4, 5), MakeIntList(), MakeIntList(6, 7, 8)).
		Flatten().
		Foreach(func(e int) { fmt.Println("flatten", e) })

	fmt.Println("reduce list", MakeIntList(10, 20, 30).Reduce(func(a, b int) int { return a + b }))
}
