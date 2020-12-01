package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp"
)

type XArr []int

func (a XArr) Foo() {}
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

	var i1 IntOption = IntOpt(1)
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

	fmt.Println(IntArray([]int{1, 2, 3, 4, 5}).MkString("~~", "|", "~~"))

	var x int = 100
	var isXdefined bool = &x == nil
	fmt.Println(isXdefined)

	mapString := MakeIntList(1, 2, 3, 4, 5).MapString(func(e int) string { return fmt.Sprintf("{%v}", e) })
	mapString.Foreach(func(e string) {
		fmt.Println("> ", e)
	})
	fmt.Println(mapString.ToString())

	fmt.Println(IntOpt(10).ToString())

	fmt.Println(MakeIntList(1, 2).Tail().Tail().HeadOption().ToString())

	fmt.Println(MakeIntList(1, 2, 3).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntList(1, 2).Equals(MakeIntList(1, 2, 3)))
	fmt.Println(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5)).
		Equals(MakeIntListList(MakeIntList(1, 2), MakeIntList(3), MakeIntList(4, 5))))

	MakeIntListList(MakeIntList(1, 2, 3), MakeIntList(4, 5), MakeIntList(), MakeIntList(6, 7, 8)).
		FlatMapInt(func(l IntList) IntList { return l }).
		Foreach(func(e int) { fmt.Println(">", e) })

	MakeIntOptionOption(IntOpt(10)).FlatMapInt(func(e IntOption) IntOption { return e }).Foreach(func(e int) { fmt.Println("flatten", e) })

	MakeIntListList(MakeIntList(1, 2, 3), MakeIntList(4, 5), MakeIntList(), MakeIntList(6, 7, 8)).
		Flatten().
		Foreach(func(e int) { fmt.Println("flatten", e) })

	fmt.Println("reduce list", MakeIntList(10, 20, 30).Reduce(func(a, b int) int { return a + b }))

	var ii int = Int(10015).Underlined()
	fmt.Println(ii)

	var parsed Int = String("10015").ToInt()
	fmt.Println("parsed", parsed, parsed.Underlined())

	fmt.Println(String("10015").ToIntOption().ToString())
	fmt.Println(String("Hello").ToIntOption().ToString())

	fmt.Println(Int(1).Cons(2).Cons(3).Cons(4).Cons(5).ToString())

	fmt.Println(Int(10).Min(20), Int(10).Min(10), Int(20).Min(10))
	fmt.Println(Int(10).Max(20), Int(10).Max(10), Int(20).Max(10))

	fmt.Println(Int(0).To(10).ToString())
	fmt.Println(Int(0).Until(10).ToString())

	xsForSearch := MakeIntList(1, -2, 3, 4, -5, 6, -7, 8, 9)
	fmt.Println(xsForSearch.Find(func(e int) bool { return e == 3 }).ToString())
	fmt.Println(xsForSearch.Find(func(e int) bool { return e == -3 }).ToString())

	xsForGroup := MakeIntList(1, 2, 3, 1, 2, 3, 1, 2, 3, 1, 1, 3, 3, 4, 5)
	var groups map[int]IntList = xsForGroup.GroupByInt(IntIdentity)
	for k, v := range groups {
		fmt.Println(k, v.ToString())
	}

	type Shape struct {
		name  string
		color string
		area  int
	}

	shapes := NilAny.
		Cons(Shape{"circle", "green", 10}).
		Cons(Shape{"circle", "red", 20}).
		Cons(Shape{"circle", "yellow", 30}).
		Cons(Shape{"triangle", "green", 10}).
		Cons(Shape{"triangle", "red", 20}).
		Cons(Shape{"triangle", "yellow", 30}).
		Cons(Shape{"polygon", "green", 10}).
		Cons(Shape{"polygon", "red", 20}).
		Cons(Shape{"polygon", "yellow", 30})

	byShape := func(e Any) Any { return e.(Shape).name }
	byColor := func(e Any) Any { return e.(Shape).color }
	byArea := func(e Any) Any { return e.(Shape).area }

	fmt.Println("by shape")
	for k, v := range shapes.GroupByAny(byShape) {
		fmt.Println(k, v.ToString())
	}
	fmt.Println("by color")
	for k, v := range shapes.GroupByAny(byColor) {
		fmt.Println(k, v.ToString())
	}
	fmt.Println("by area")
	for k, v := range shapes.GroupByAny(byArea) {
		fmt.Println(k, v.ToString())
	}

	a2 := IntArray([]int{1, 2, 3})
	fmt.Println(a2.Head(), a2.HeadOption().ToString(), a2.Tail().MkString("[", "|", "]"))
	fmt.Println(IntArray([]int{1, 2, 3, 4, 5}).Filter(func(e int) bool { return e%2 == 0 }).ToString())
}
