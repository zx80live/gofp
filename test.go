package main

import (
	"fmt"
	. "github.com/zx80live/gofp/fp/lazy"
)

//type LazyState = func() State
//
//type State struct {
//	head *LazyInt
//	tail *LazyList
//}
//
//var EmptyState State = State{nil, nil}
//var NilLazyList LazyList = LazyList{nil}
//
//type LazyList struct{ state *LazyState }
//
//func (l LazyList) IsEmpty() bool { return l == NilLazyList }
//
//func (l LazyList) NonEmpty() bool { return !l.IsEmpty() }
//
//func (l LazyList) Cons(i LazyInt) LazyList {
//	xs := LazyList{l.state}
//	state := func() State { return State{&i, &xs} }
//
//	return LazyList{&state}
//}
//
//func (l LazyList) Head() (int, LazyList) {
//	if l.IsEmpty() {
//		panic("can't get head from empty list")
//	} else {
//		s := (*l.state)()
//		return s.head.Eval().Value(), *s.tail
//	}
//}
//
//func (l LazyList) HeadOption() (fp.IntOption, LazyList) {
//	if l.IsEmpty() {
//		return fp.NoneInt, NilLazyList
//	} else {
//		s := (*l.state)()
//		return fp.MkIntOption(s.head.Eval().Value()), *s.tail
//	}
//}
//
//func (l LazyList) Map(f func(e int) int) LazyList {
//	newState := func() State {
//		state := (*l.state)()
//		h := *state.head
//		mappedValue := func() int {
//			return f(h.Value())
//		}
//		mappedH := MkLazyInt(mappedValue)
//		t := state.tail.Map(f)
//
//		return State{&mappedH, &t}
//	}
//	return LazyList{&newState}
//}
//
//func (l LazyList) Filter(p func(e int) bool) LazyList {
//	newState := func() State {
//		var xs = l
//		var h LazyInt
//		var found = false
//
//		for !found && xs.NonEmpty() {
//			s := (*xs.state)()
//			h = (*s.head).Eval()
//			found = p(h.Value())
//			xs = *s.tail
//		}
//
//		if found {
//			t := xs.Filter(p)
//			return State{&h, &t}
//		} else {
//			return EmptyState
//		}
//	}
//
//	return LazyList{&newState}
//}
//
//func (l LazyList) Take(n int) LazyList {
//	if n <= 0 {
//		return NilLazyList
//	}
//
//	newState := func() State {
//		state := (*l.state)()
//		t := state.tail.Take(n - 1)
//		return State{state.head, &t}
//	}
//	return LazyList{&newState}
//}

// Creates infinite lazy list of numbers
func InfiniteIntList(from int) IntLazyList {
	el := MkLazyInt(func() int { return from })
	var state IntLazyState = func() IntState {
		rest := InfiniteIntList(from + 1)
		return MkIntState(el, rest)
	}
	return MkIntLazyList(state)
}

// Creates finite lazy list of numbers
//func FiniteIntList(from, to int) LazyList {
//	if from == to {
//		return NilLazyList
//	}
//
//	el := MkLazyInt(func() int { return from })
//
//	state := func() State {
//		rest := FiniteIntList(from+1, to)
//		return State{&el, &rest}
//	}
//	return LazyList{&state}
//}

func main() {
	fmt.Println("Hello Lazy")

	xs := InfiniteIntList(1).
		//Map(func(e int) int { return e + 1 }).
		//Map(func(e int) int { return e - 1 }).
		Filter(func(e int) bool { return e%2 == 0 }).
		Filter(func(e int) bool { return e > -100 }).
		Filter(func(e int) bool { return e > 0 }).
		//Map(func(e int) int { return e + 1 }).
		Filter(func(e int) bool { return e > 10 }).Take(5)

	for i := 0; xs.NonEmpty() && i < 100; i++ {
		h, t := xs.Head()

		fmt.Println(h)
		xs = t
	}
}
