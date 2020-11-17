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
	var i2 IntOption = NoneIntOption

	fmt.Println(i1.ToString(), i1.IsDefined(), i1.IsEmpty(), i1.Filter(even).ToString(), i1.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())
	fmt.Println(i2.ToString(), i2.IsDefined(), i2.IsEmpty(), i2.Filter(even).ToString(), i2.Filter(pos).ToString(), i1.MapInt(func(e int) int { return e + 100 }).ToString())

	fmt.Println(NoneIntOption.Equals(NoneIntOption))

	fmt.Println(MakeIntList(1, 2, 3, 4, 5).Size())
	fmt.Println(MakeIntList(1, 2).Size())
	fmt.Println(MakeIntList(1).Size())
	fmt.Println(NilIntList.Size())

	fmt.Println(MakeIntList(1, 2, 3).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList(1, 2, 3, 4).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList().Equals(MakeIntList()))

	MakeIntList(1, 2, 3).MapInt(func(e int) int { return e * 100 }).Foreach(func(e int) { fmt.Println(">", e) })

	arr := MakeIntList(1, 2, 3, 4, 5).ToArray()
	fmt.Println(len(arr), arr[0], arr[1], arr[2], arr[3], arr[4])

	fmt.Println(MkStringIntArr([]int{1, 2, 3, 4, 5}, "~~", "|", "~~"))
	nestedList := MakeIntArrList([]int{1, 2, 3}, []int{4, 5, 6}, []int{7})
	fmt.Println(MkStringIntArrList(nestedList, "[", ",", "]"))
	fmt.Println(nestedList.ToString())

	var x int = 100
	var isXdefined bool = &x == nil
	fmt.Println(isXdefined)

	mapString := MakeIntList(1, 2, 3, 4, 5).MapString(func(e int) string { return fmt.Sprintf("{%v}", e) })
	mapString.Foreach(func(e string) {
		fmt.Println("> ", e)
	})
	fmt.Println(mapString.ToString())

	sum1 := Int(10).FlatMapInt(func(i int) IntOption {
		return Int(20).MapInt(func(j int) int {
			return i + j
		})
	})

	fmt.Println(sum1.ToString())

	xs1 := MakeIntListList(MakeIntList(1, 2), NilIntList, MakeIntList(3), MakeIntList(4, 5, 6))
	fmt.Println(xs1.ToString())
	res1 := xs1.FlatMapInt(func(list IntList) IntList {
		fmt.Println(" flatMap.iterate ->", list.ToString())
		return list
	})

	fmt.Println(res1.ToString())

	MakeIntList(1, 2, 3, 4, 5).
		FlatMapInt(func(i int) IntList { return MakeIntList(i * 10) }).
		Foreach(func(i int) { fmt.Println(" ", i) })

	fmt.Println(MakeIntList(1, 2).Tail().Tail().HeadOption().ToString())

	fmt.Println(MakeIntList(1, 2, 3).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList(1, 2).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5)).
		Equals(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5))))
}
